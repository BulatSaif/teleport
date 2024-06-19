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

package installer

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gravitational/trace"

	"github.com/gravitational/teleport/lib/defaults"
)

type packageManagerYUM struct {
	*packageManagerYUMConfig
}

type packageManagerYUMConfig struct {
	httpClient *http.Client

	bins binariesLocation

	// fsRootPrefix is the prefix to use when reading operating system information and when installing teleport.
	// Used for testing.
	fsRootPrefix string

	logger *slog.Logger
}

func (p *packageManagerYUMConfig) checkAndSetDefaults() error {
	var err error
	if p == nil {
		return trace.BadParameter("config is required")
	}

	if p.httpClient == nil {
		p.httpClient, err = defaults.HTTPClient()
		if err != nil {
			return trace.Wrap(err)
		}
	}

	p.bins.checkAndSetDefaults()

	if p.fsRootPrefix == "" {
		p.fsRootPrefix = "/"
	}

	if p.logger == nil {
		p.logger = slog.Default()
	}

	return nil
}

// newPackageManagerYUM creates a new PackageManagerYUM.
func newPackageManagerYUM(cfg *packageManagerYUMConfig) (*packageManagerYUM, error) {
	if err := cfg.checkAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}
	return &packageManagerYUM{packageManagerYUMConfig: cfg}, nil
}

const yumRepoEndpoint = "https://yum.releases.teleport.dev/"

// AddTeleportRepository adds the Teleport repository to the current system.
func (pm *packageManagerYUM) AddTeleportRepository(ctx context.Context, linuxInfo *linuxDistroInfo, repoChannel string) error {
	// Teleport repo only targets the major version of the target distros.
	versionID := strings.Split(linuxInfo.VersionID, ".")[0]

	pm.logger.InfoContext(ctx, "Installing yum-utils", "command", "yum install -y yum-utils")
	installYumUtilsCMD := exec.CommandContext(ctx, pm.bins.yum, "install", "-y", "yum-utils")
	installYumUtilsCMDOutput, err := installYumUtilsCMD.CombinedOutput()
	if err != nil {
		return trace.Wrap(err, string(installYumUtilsCMDOutput))
	}

	// Repo location looks like this:
	// https://yum.releases.teleport.dev/$ID/$VERSION_ID/Teleport/%{_arch}/{{ .RepoChannel }}/teleport.repo
	repoLocation := fmt.Sprintf(`%s%s/%s/Teleport/%%{_arch}/%s/teleport.repo`, yumRepoEndpoint, linuxInfo.ID, versionID, repoChannel)
	pm.logger.InfoContext(ctx, "Building rpm metadata for Teleport repo", "command", "rpm --eval "+repoLocation)
	rpmEvalTeleportRepoCMD := exec.CommandContext(ctx, pm.bins.rpm, "--eval", repoLocation)
	rpmEvalTeleportRepoCMDOutput, err := rpmEvalTeleportRepoCMD.CombinedOutput()
	if err != nil {
		return trace.Wrap(err, string(rpmEvalTeleportRepoCMDOutput))
	}

	// output from the command above might have a `\n` at the end.
	repoURL := strings.TrimSpace(string(rpmEvalTeleportRepoCMDOutput))

	pm.logger.InfoContext(ctx, "Adding repository metadata", "command", "yum-config-manager --add-repo "+repoURL)
	yumAddRepoCMD := exec.CommandContext(ctx, pm.bins.yumConfigManager, "--add-repo", repoURL)
	yumAddRepoCMDOutput, err := yumAddRepoCMD.CombinedOutput()
	if err != nil {
		return trace.Wrap(err, string(yumAddRepoCMDOutput))
	}

	return nil
}

// InstallPackages installs one or multiple packages into the current system.
func (pm *packageManagerYUM) InstallPackages(ctx context.Context, packageList []packageVersion) error {
	if len(packageList) == 0 {
		return nil
	}

	installArgs := make([]string, 0, len(packageList)+2)
	installArgs = append(installArgs, "install", "-y")

	for _, pv := range packageList {
		if pv.Version != "" {
			installArgs = append(installArgs, pv.Name+"-"+pv.Version)
			continue
		}
		installArgs = append(installArgs, pv.Name)
	}

	pm.logger.InfoContext(ctx, "Installing", "command", "yum "+strings.Join(installArgs, " "))

	installPackagesCMD := exec.CommandContext(ctx, pm.bins.yum, installArgs...)
	installPackagesCMDOutput, err := installPackagesCMD.CombinedOutput()
	if err != nil {
		return trace.Wrap(err, string(installPackagesCMDOutput))
	}
	return nil
}
