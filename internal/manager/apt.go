package manager

import (
	"log"
	"os/exec"
)

type Apt struct {
}

// InstallPackage install a package with the apt package manager
func (a Apt) InstallPackage(name string) (bool, error) {
	cmd := exec.Command("apt", "install", name, "-y")
	log.Println("Installing package", name)
	if ok, err := runCommand(cmd); !ok {
		log.Println("Could not install the package ", err)
		return false, err
	}
	log.Println("Package installed", name)
	return true, nil
}

// UninstallPackage uninstall a package with the apt package manager
func (a Apt) UninstallPackage(name string) (bool, error) {
	cmd := exec.Command("apt", "remove", name, "-y")
	log.Println("Uninstalling package", name)
	if ok, err := runCommand(cmd); !ok {
		log.Println("Could not uninstall the package ", err)
		return false, err
	}
	log.Println("Package uninstalled", name)
	return true, nil
}
