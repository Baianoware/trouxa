package aptitude

import (
	"os/exec"
)

type Commander struct {
}

// BuildInstallCommand create the installation command to aptitude
func (a *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("aptitude", "install", "-y", name)
}

// BuildUninstallCommand create the uninstallation command to aptitude
func (a *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("aptitude", "remove", "-y", name)
}
