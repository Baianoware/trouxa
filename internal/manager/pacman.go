package manager

import (
	"log"
	"os/exec"
)

type Pacman struct {
}

// InstallPackage install a package with the Pacman package manager
func (a Pacman) InstallPackage(name string) (bool, error) {
	cmd := exec.Command("pacman", "-Sy", name, "--noconfirm")
	log.Println("Installing package", name)
	if ok, err := runCommand(cmd); !ok {
		log.Println("Could not install the package ", err)
		return false, err
	}
	log.Println("Package installed", name)
	return true, nil
}

// UninstallPackage uninstall a package with the Pacman package manager
func (a Pacman) UninstallPackage(name string) (bool, error) {
	cmd := exec.Command("pacman", "-Rs", name, "--noconfirm")
	log.Println("Uninstalling package", name)
	if ok, err := runCommand(cmd); !ok {
		log.Println("Could not uninstall the package ", err)
		return false, err
	}
	log.Println("Package uninstalled", name)
	return true, nil
}
