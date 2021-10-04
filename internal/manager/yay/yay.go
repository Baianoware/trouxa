package yay

import (
	"os/exec"
)

type Commander struct {
}

// BuildInstallCommand create the installation command to yay
func (p *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("yay", "-Sy", name, "--noconfirm")
}

// BuildUninstallCommand create the uninstallation command to yay
func (p *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("yay", "-Rs", name, "--noconfirm")
}
