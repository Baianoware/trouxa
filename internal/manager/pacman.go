package manager

import (
	"log"
	"os/exec"
)

type Pacman struct {
}

// InstallPackage install a package with the Pacman package manager
func (a Pacman) InstallPackage(name string) (error, bool) {
	cmd := exec.Command("pacman", "-Sy", name, "--noconfirm")
	log.Println("Installing package", name)
	if err, ok := runCommand(cmd); !ok {
		log.Println("Could not install the package ", err)
		return err, false
	}
	log.Println("Package installed", name)
	return nil, true
}

// UninstallPackage uninstall a package with the Pacman package manager
func (a Pacman) UninstallPackage(name string) (error, bool) {
	cmd := exec.Command("pacman", "-Rs", name, "--noconfirm")
	log.Println("Uninstalling package", name)
	if err, ok := runCommand(cmd); !ok {
		log.Println("Could not uninstall the package ", err)
		return err, false
	}
	log.Println("Package uninstalled", name)
	return nil, true
}
