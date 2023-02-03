package commander

import (
	"fmt"
	"os/exec"

	"github.com/Bainoware/trouxa/internal/manager"
	"github.com/Bainoware/trouxa/internal/manager/apk"
	"github.com/Bainoware/trouxa/internal/manager/apt"
	"github.com/Bainoware/trouxa/internal/manager/aptitude"
	"github.com/Bainoware/trouxa/internal/manager/dnf"
	"github.com/Bainoware/trouxa/internal/manager/eopkg"
	"github.com/Bainoware/trouxa/internal/manager/pacman"
	"github.com/Bainoware/trouxa/internal/manager/snap"
	"github.com/Bainoware/trouxa/internal/manager/yay"
	"github.com/Bainoware/trouxa/internal/manager/yum"
	"github.com/Bainoware/trouxa/internal/manager/zypper"
)

// Commander interface with methods used to build install and uninstall commands
type Commander interface {
	BuildInstallCommand(name string) *exec.Cmd
	BuildUninstallCommand(name string) *exec.Cmd
	ListPackages() *exec.Cmd
}

// FromName create a commander from a package manager mane
func FromName(name string) (Commander, error) {
	if err := manager.IsValid(name); err != nil {
		return nil, fmt.Errorf("Package manager not found on user's path: %w", err)
	}
	// To no use reflection...
	switch name {
	case "apt":
		return new(apt.Commander), nil
	case "apk":
		return new(apk.Commander), nil
	case "aptitude":
		return new(aptitude.Commander), nil
	case "dnf":
		return new(dnf.Commander), nil
	case "eopkg":
		return new(eopkg.Commander), nil
	case "pacman":
		return new(pacman.Commander), nil
	case "snap":
		return new(snap.Commander), nil
	case "yay":
		return new(yay.Commander), nil
	case "yum":
		return new(yum.Commander), nil
	case "zypper":
		return new(zypper.Commander), nil
	default:
		return nil, fmt.Errorf("Package manager not found on user's path")
	}
}
