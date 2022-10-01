package snap

import (
	"os/exec"
)

type Commander struct {
}

// BuildInstallCommand create the installation command to snap
func (p *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("snap", "install", name)
}

// BuildUninstallCommand create the uninstallation command to snap
func (p *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("snap", "remove", name)
}
