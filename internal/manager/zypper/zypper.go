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

// ListPackages lists all installed packages from zypper
func (a *Commander) ListPackages() *exec.Cmd {
	return exec.Command("zypper", "search", "--installed-only")
}
