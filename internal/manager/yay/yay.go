package yay

import (
	"os/exec"
)

type Commander struct {
}

func (p *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("yay", "-Sy", name, "--noconfirm")
}
func (p *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("yay", "-Rs", name, "--noconfirm")
}
