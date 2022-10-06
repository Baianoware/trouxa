package dnf

import (
	"os/exec"
)

type Commander struct {
}

// BuildInstallCommand create the installation command to dnf
func (p *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("dnf", "install", "-y", name)
}

// BuildUninstallCommand create the uninstallation command to dnf
func (p *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("dnf", "remove", "-y", name)
}

// ListPackages lists all installed packages from dnf
func (a *Commander) ListPackages() *exec.Cmd {
	return exec.Command("dnf", "list", "installed")
}
