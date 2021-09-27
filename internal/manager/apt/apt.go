package apt

import (
	"os/exec"
)

type Commander struct {
}

func (a *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("apt", "install", name, "-y")
}
func (a *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("apt", "remove", name, "-y")
}
