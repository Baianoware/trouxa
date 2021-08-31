package manager

import (
	"log"
	"os/exec"
)

type Apt struct {
}

// InstallPackage install a package with the apt package manager
func (a Apt) InstallPackage(name string) (error, bool) {
	cmd := exec.Command("apt", "install", name, "-y")
	log.Println("Installing package", name)
	if err, ok := runCommand(cmd); !ok {
		log.Println("Could not install the package ", err)
		return err, false
	}
	log.Println("Package installed", name)
	return nil, true
}

// UninstallPackage uninstall a package with the apt package manager
func (a Apt) UninstallPackage(name string) (error, bool) {
	cmd := exec.Command("apt", "remove", name, "-y")
	log.Println("Uninstalling package", name)
	if err, ok := runCommand(cmd); !ok {
		log.Println("Could not uninstall the package ", err)
		return err, false
	}
	log.Println("Package uninstalled", name)
	return nil, true
}
