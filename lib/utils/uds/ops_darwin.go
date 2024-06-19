//go:build darwin && cgo

// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package uds

// these functions are internal macos-only libpthread functions that the
// Chromium team felt comfortable enough to use (at the time of writing this
// comment,
// https://chromium.googlesource.com/chromium/src/base/+/50531cb2906c6b2a69ee0a9d10783760cb3922fc/process/launch_mac.cc
// ); the functions are documented at
// https://opensource.apple.com/source/libpthread/libpthread-454.80.2/private/pthread/private.h.auto.html

// int pthread_chdir_np(const char *dir);
// int pthread_fchdir_np(int fd);
import "C"

import (
	"context"
	"net"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"unsafe"

	"github.com/gravitational/trace"
)

const sunPathLen = int(unsafe.Sizeof(syscall.RawSockaddrUnix{}.Path))

// ListenUnix is like [net.ListenUnix] but with a context (or like
// [net.ListenConfig.Listen] without a type assertion). The network must be
// "unix" or "unixpacket". On this platform (darwin), only the last component of
// the path must fit in sockaddr_un (104 characters), as oversized paths are
// handled by changing directory before binding the socket.
func (lc *ListenConfig) ListenUnix(ctx context.Context, network, path string) (*net.UnixListener, error) {
	switch network {
	case "unix", "unixpacket":
	default:
		return nil, trace.BadParameter("invalid network %q, expected \"unix\" or \"unixpacket\"", network)
	}

	if strings.IndexByte(path, '\x00') != -1 {
		return nil, trace.BadParameter("path must not contain NUL")
	}

	if len(path) > sunPathLen {
		path = filepath.Clean(path)
	}

	if len(path) <= sunPathLen {
		l, err := (*net.ListenConfig)(lc).Listen(ctx, network, path)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		return l.(*net.UnixListener), nil
	}

	dir, file := filepath.Split(path)
	if len(file) > sunPathLen {
		return nil, trace.BadParameter("final path component is too long")
	}

	dir = filepath.Clean(dir) + "\x00"
	var bDir *byte = unsafe.StringData(dir)
	cDir := (*C.char)(unsafe.Pointer(bDir))

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	r, errno := C.pthread_chdir_np(cDir)
	if r != 0 {
		return nil, trace.Wrap(trace.ConvertSystemError(errno))
	}
	defer C.pthread_fchdir_np(-1)

	l, err := (*net.ListenConfig)(lc).Listen(ctx, network, file)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return l.(*net.UnixListener), err
}

// ListenUnixgram is like [net.ListenUnixgram] but with a context (or like
// [net.ListenConfig.ListenPacket] without a type assertion). The network must
// be "unixgram". On this platform (darwin), only the last component of the path
// must fit in sockaddr_un (104 characters), as oversized paths are handled by
// changing directory before binding the socket.
func (lc *ListenConfig) ListenUnixgram(ctx context.Context, network, path string) (*net.UnixConn, error) {
	switch network {
	case "unixgram":
	default:
		return nil, trace.BadParameter("invalid network %q, expected \"unixgram\"", network)
	}

	if strings.IndexByte(path, '\x00') != -1 {
		return nil, trace.BadParameter("path must not contain NUL")
	}

	if len(path) > sunPathLen {
		path = filepath.Clean(path)
	}

	if len(path) <= sunPathLen {
		l, err := (*net.ListenConfig)(lc).ListenPacket(ctx, network, path)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		return l.(*net.UnixConn), nil
	}

	dir, file := filepath.Split(path)
	if len(file) > sunPathLen {
		return nil, trace.BadParameter("final path component is too long")
	}

	dir = filepath.Clean(dir) + "\x00"
	var bDir *byte = unsafe.StringData(dir)
	cDir := (*C.char)(unsafe.Pointer(bDir))

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	r, errno := C.pthread_chdir_np(cDir)
	if r != 0 {
		return nil, trace.Wrap(trace.ConvertSystemError(errno))
	}
	defer C.pthread_fchdir_np(-1)

	l, err := (*net.ListenConfig)(lc).ListenPacket(ctx, network, file)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return l.(*net.UnixConn), err
}

// DialUnix is like [net.DialUnix] but with a context (or like
// [net.Dialer.DialContext] without a type assertion). The network must be
// "unix", "unixgram" or "unixpacket". On this platform (darwin), only the last
// component of the path must fit in sockaddr_un (104 characters), as oversized
// paths are handled by changing directory before binding the socket.
func (d *Dialer) DialUnix(ctx context.Context, network, path string) (*net.UnixConn, error) {
	switch network {
	case "unix", "unixgram", "unixpacket":
	default:
		return nil, trace.BadParameter("invalid network %q, expected \"unix\", \"unixgram\" or \"unixpacket\"", network)
	}

	if strings.IndexByte(path, '\x00') != -1 {
		return nil, trace.BadParameter("path must not contain NUL")
	}

	if len(path) > sunPathLen {
		path = filepath.Clean(path)
	}

	if len(path) <= sunPathLen {
		l, err := (*net.Dialer)(d).DialContext(ctx, network, path)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		return l.(*net.UnixConn), nil
	}

	dir, file := filepath.Split(path)
	if len(file) > sunPathLen {
		return nil, trace.BadParameter("final path component is too long")
	}

	dir = filepath.Clean(dir) + "\x00"
	cDir := (*C.char)(unsafe.Pointer(unsafe.StringData(dir)))

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	r, errno := C.pthread_chdir_np(cDir)
	if r != 0 {
		return nil, trace.Wrap(trace.ConvertSystemError(errno))
	}
	defer C.pthread_fchdir_np(-1)

	l, err := (*net.Dialer)(d).DialContext(ctx, network, file)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return l.(*net.UnixConn), err
}
