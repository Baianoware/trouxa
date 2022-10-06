package yum

import (
	"os/exec"
)

type Commander struct {
}

// BuildInstallCommand create the installation command to yum
func (p *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("yum", "install", "-y", name)
}

// BuildUninstallCommand create the uninstallation command to yum
func (p *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("yum", "remove", "-y", name)
}

// ListPackages lists all installed packages from yum
func (a *Commander) ListPackages() *exec.Cmd {
	return exec.Command("yum", "list", "installed")
}
