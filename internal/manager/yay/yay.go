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

// DumpPackages lists all installed packages from yay
func (a *Commander) DumpPackages() *exec.Cmd {
	// There seems to be no way of distinguishing between packages installed by pacman and yay:
	// https://www.reddit.com/r/archlinux/comments/woh8fr/comment/ikb075w/?utm_source=reddit&utm_medium=web2x&context=3
	// So the command ends up being virtually the same here
	return exec.Command("yay", "-Qe")
}
