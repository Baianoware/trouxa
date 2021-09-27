package pacman

import (
	"os/exec"
)

type Commander struct {
}

func (p *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("pacman", "-Sy", name, "--noconfirm")
}
func (p *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("pacman", "-Rs", name, "--noconfirm")
}
