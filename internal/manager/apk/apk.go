package apk

import (
	"os/exec"
)

type Commander struct {
}

func (p *Commander) BuildInstallCommand(name string) *exec.Cmd {
	return exec.Command("apk", "add", name)
}
func (p *Commander) BuildUninstallCommand(name string) *exec.Cmd {
	return exec.Command("apk", "del", name)
}
