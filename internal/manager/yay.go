package manager

import (
	"log"
	"os/exec"
)

type Yay struct {
}

// InstallPackage install a package with the yay package manager
func (a Yay) InstallPackage(name string) (error, bool) {
	cmd := exec.Command("yay", "-Sy", name, "--noconfirm")
	log.Println("Installing package", name)
	if err, ok := runCommand(cmd); !ok {
		log.Println("Could not install the package ", err)
		return err, false
	}
	log.Println("Package installed", name)
	return nil, true
}

// UninstallPackage uninstall a package with the yay package manager
func (a Yay) UninstallPackage(name string) (error, bool) {
	cmd := exec.Command("yay", "-Rs", name, "--noconfirm")
	log.Println("Uninstalling package", name)
	if err, ok := runCommand(cmd); !ok {
		log.Println("Could not uninstall the package ", err)
		return err, false
	}
	log.Println("Package uninstalled", name)
	return nil, true
}
