package apk

import (
	"os/exec"
)

type Commander struct{}

// BuildInstallCommand create the installation command to apk
func (p *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("apk", "add", name)
}

// BuildUninstallCommand create the uninstallation command to apk
func (p *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("apk", "del", name)
}

// ListPackages lists all installed packages from apk
func (a *Commander) ListPackages() *exec.Cmd {
	return exec.Command("apk", "list", "--installed")
}
