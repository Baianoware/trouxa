package pacman

import (
	"os/exec"
)

type Commander struct {
}

// BuildInstallCommand create the installation command to pacman
func (p *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("pacman", "-Sy", name, "--noconfirm")
}

// BuildUninstallCommand create the uninstallation command to pacman
func (p *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("pacman", "-Rs", name, "--noconfirm")
}

// DumpPackages lists all installed packages from pacman
func (a *Commander) DumpPackages() *exec.Cmd {
	return exec.Command("pacman", "-Qe")
}
