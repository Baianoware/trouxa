package commander

import (
	"os/exec"

	"github.com/Bainoware/trouxa/internal/manager/apt"
	"github.com/Bainoware/trouxa/internal/manager/pacman"
	"github.com/Bainoware/trouxa/internal/manager/yay"
)

// Commander interface with methods used to build install and unistall commands
type Commander interface {
	BuildInstallCommand(name string) *exec.Cmd
	BuildUninstallCommand(name string) *exec.Cmd
}

// FromName create a commander from a package manager mane
func FromName(manager string) Commander {
	// To no use reflection...
	switch manager {
	case "apt":
		return new(apt.Commander)
	case "pacman":
		return new(pacman.Commander)
	case "yay":
		return new(yay.Commander)
	default:
		return nil
	}
}
