/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package tbot

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"os"
	"runtime"
	"time"

	"github.com/gravitational/trace"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gravitational/teleport"
	machineidv1pb "github.com/gravitational/teleport/api/gen/proto/go/teleport/machineid/v1"
	"github.com/gravitational/teleport/api/utils/retryutils"
	"github.com/gravitational/teleport/lib/tbot/config"
)

type heartbeatSubmitter interface {
	SubmitHeartbeat(
		ctx context.Context, in *machineidv1pb.SubmitHeartbeatRequest, opts ...grpc.CallOption,
	) (*machineidv1pb.SubmitHeartbeatResponse, error)
}

type heartbeatService struct {
	now                func() time.Time
	log                *slog.Logger
	botCfg             *config.BotConfig
	startedAt          time.Time
	heartbeatSubmitter heartbeatSubmitter
	interval           time.Duration
	retryLimit         int
}

func (s *heartbeatService) heartbeat(ctx context.Context, isStartup bool) error {
	s.log.DebugContext(ctx, "Sending heartbeat")
	hostName, err := os.Hostname()
	if err != nil {
		s.log.WarnContext(ctx, "Failed to determine hostname for heartbeat", "error", err)
	}

	hb := &machineidv1pb.BotInstanceStatusHeartbeat{
		RecordedAt:   timestamppb.New(s.now()),
		Hostname:     hostName,
		IsStartup:    isStartup,
		Uptime:       durationpb.New(s.now().Sub(s.startedAt)),
		OneShot:      s.botCfg.Oneshot,
		JoinMethod:   string(s.botCfg.Onboarding.JoinMethod),
		Version:      teleport.Version,
		Architecture: runtime.GOARCH,
		Os:           runtime.GOOS,
	}

	_, err = s.heartbeatSubmitter.SubmitHeartbeat(ctx, &machineidv1pb.SubmitHeartbeatRequest{
		Heartbeat: hb,
	})
	if err != nil {
		return trace.Wrap(err, "submitting heartbeat")
	}

	s.log.InfoContext(ctx, "Sent heartbeat", "data", hb.String())
	return nil
}

func (s *heartbeatService) OneShot(ctx context.Context) error {
	err := s.heartbeat(ctx, true)
	// Ignore not implemented as this is likely confusing.
	// TODO(noah): Remove NotImplemented check at V18 assuming V17 first major
	// with heartbeating.
	if err != nil && !trace.IsNotImplemented(err) {
		return trace.Wrap(err)
	}
	return nil
}

func (s *heartbeatService) Run(ctx context.Context) error {
	ticker := time.NewTicker(s.interval)
	jitter := retryutils.NewJitter()
	defer ticker.Stop()

	isStartup := true
	for {
		var err error
		for attempt := 1; attempt <= s.retryLimit; attempt++ {
			s.log.InfoContext(
				ctx,
				"Attempting to send heartbeat",
				"attempt", attempt,
				"retry_limit", s.retryLimit,
			)
			err = s.heartbeat(ctx, isStartup)
			if err == nil {
				isStartup = false
				break
			}

			// If the cluster does not support Bot Instance heartbeats, we
			// do not retry.
			// TODO(noah): Remove NotImplemented check at V18 assuming V17 first
			// major with heartbeating.
			if trace.IsNotImplemented(err) {
				s.log.WarnContext(
					ctx,
					"Cluster does not support Bot Instance heartbeats. Will not attempt to send further heartbeats",
				)
				return nil
			}

			if attempt != s.retryLimit {
				// exponentially back off with jitter, starting at 1 second.
				backoffTime := time.Second * time.Duration(math.Pow(2, float64(attempt-1)))
				backoffTime = jitter(backoffTime)
				s.log.WarnContext(
					ctx,
					"Sending heartbeat failed. Waiting to retry",
					"attempt", attempt,
					"retry_limit", s.retryLimit,
					"backoff", backoffTime,
					"error", err,
				)
				select {
				case <-ctx.Done():
					return nil
				case <-time.After(backoffTime):
				}
			}
		}
		if err != nil {
			s.log.WarnContext(
				ctx,
				"All retry attempts exhausted sending heartbeat. Waiting for next heartbeat",
				"retry_limit", s.retryLimit,
				"interval", s.interval,
			)
		} else {
			s.log.InfoContext(
				ctx,
				"Sent heartbeat. Waiting for next heartbeat",
				"interval", s.interval,
			)
		}

		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			continue
		}
	}
}

func (s *heartbeatService) String() string {
	return fmt.Sprintf("heartbeat")
}
