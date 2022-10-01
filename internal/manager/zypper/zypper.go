package zypper

import (
	"os/exec"
)

type Commander struct {
}

// BuildInstallCommand create the installation command to zypper
func (p *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("zypper", "install", "-y", name)
}

// BuildUninstallCommand create the uninstallation command to zypper
func (p *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("zypper", "remove", "-y", name)
}
