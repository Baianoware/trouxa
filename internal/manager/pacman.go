package manager

import (
	"log"
	"os/exec"
)

//TODO Fix it
type ManagerPacman struct {
}

// InstallPackage install a package with the Pacman package manager
func (a ManagerPacman) InstallPackage(name string) bool {
	cmd := exec.Command("pacman", "-Sy", name, "--noconfirm")
	log.Println("Installing package", name)
	if ok := runCommand(cmd); !ok {
		log.Fatalf("Could not install the package")
		return false
	}
	log.Println("Package installed", name)
	return true
}

// UninstallPackage uninstall a package with the Pacman package manager
func (a ManagerPacman) UninstallPackage(name string) bool {
	cmd := exec.Command("pacman", "-Rs", name, "--noconfirm")
	log.Println("Uninstalling package", name)
	if ok := runCommand(cmd); !ok {
		log.Fatal("Could not uninstall the package")
		return false
	}
	log.Println("Package uninstalled", name)
	return true
}
