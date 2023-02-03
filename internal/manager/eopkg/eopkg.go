package eopkg

import (
	"os/exec"
)

type Commander struct{}

// BuildInstallCommand create the installation command to eopkg
func (p *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("eopkg", "install", "-y", name)
}

// BuildUninstallCommand create the uninstallation command to eopkg
func (p *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("eopkg", "remove", "-y", name)
}

// ListPackages lists all installed packages from eopkg
func (a *Commander) ListPackages() *exec.Cmd {
	return exec.Command("eopkg", "li")
}
