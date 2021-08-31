package manager

import (
	"log"
	"os/exec"
)

type Yay struct {
}

// InstallPackage install a package with the yay package manager
func (a Yay) InstallPackage(name string) (bool, error) {
	cmd := exec.Command("yay", "-Sy", name, "--noconfirm")
	log.Println("Installing package", name)
	if ok, err := runCommand(cmd); !ok {
		log.Println("Could not install the package ", err)
		return false, err
	}
	log.Println("Package installed", name)
	return true, nil
}

// UninstallPackage uninstall a package with the yay package manager
func (a Yay) UninstallPackage(name string) (bool, error) {
	cmd := exec.Command("yay", "-Rs", name, "--noconfirm")
	log.Println("Uninstalling package", name)
	if ok, err := runCommand(cmd); !ok {
		log.Println("Could not uninstall the package ", err)
		return false, err
	}
	log.Println("Package uninstalled", name)
	return true, nil
}
