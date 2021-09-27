package manager

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// runCommand Execute a command and send the output to default Stdout and Stderr
func runCommand(cmd *exec.Cmd) (bool, error) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Debugln("Trying to run the command")
	err := cmd.Run()
	if err != nil {
		log.Errorln("Could not run the command ", err)

		return false, err
	}
	log.Debug("Command executed")

	return true, nil
}

// InstallPackage install a package with the apt package manager
func InstallPackage(name string, cmd *exec.Cmd) (bool, error) {
	log.Infoln("Installing package: ", name)
	if ok, err := runCommand(cmd); !ok {
		log.Errorln("Could not install the package. ", err)

		return false, err
	}
	log.Infoln("Package installed: ", name)

	return true, nil
}

// UninstallPackage uninstall a package with the apt package manager
func UninstallPackage(name string, cmd *exec.Cmd) (bool, error) {
	log.Infoln("Uninstalling package: ", name)
	if ok, err := runCommand(cmd); !ok {
		log.Errorln("Could not uninstall the package ", err)

		return false, err
	}
	log.Errorln("Package uninstalled: ", name)

	return true, nil
}
