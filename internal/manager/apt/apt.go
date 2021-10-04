package apt

import (
	"os/exec"
)

type Commander struct {
}

// BuildInstallCommand create the installation command to apt
func (a *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("apt", "install", name, "-y")
}

// BuildUninstallCommand create the uninstallation command to apt
func (a *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("apt", "remove", name, "-y")
}
