package commander

import (
	"os/exec"

	"github.com/Bainoware/trouxa/internal/manager"
	"github.com/Bainoware/trouxa/internal/manager/apk"
	"github.com/Bainoware/trouxa/internal/manager/apt"
	"github.com/Bainoware/trouxa/internal/manager/pacman"
	"github.com/Bainoware/trouxa/internal/manager/yay"
	log "github.com/sirupsen/logrus"
)

// Commander interface with methods used to build install and uninstall commands
type Commander interface {
	BuildInstallCommand(name string) *exec.Cmd
	BuildUninstallCommand(name string) *exec.Cmd
}

// FromName create a commander from a package manager mane
func FromName(name string) Commander {
	if err := manager.IsValid(name); err != nil {
		log.Debugln("Package manager not found on user's path")
		return nil
	}
	// To no use reflection...
	switch name {
	case "apt":
		return new(apt.Commander)
	case "apk":
		return new(apk.Commander)
	case "pacman":
		return new(pacman.Commander)
	case "yay":
		return new(yay.Commander)
	default:
		return nil
	}
}
