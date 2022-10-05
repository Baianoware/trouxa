package manager

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func runCommand(cmd *exec.Cmd) (bool, error) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Debugln("Trying to run the command")
	err := cmd.Run()
	if err != nil {
		log.WithError(err).Debugln("Could not run the command")

		return false, err
	}
	log.Debug("Command executed")

	return true, nil
}

// IsValid check if the package manager is set on user's PATH
func IsValid(name string) error {
	_, err := exec.LookPath(name)
	if err != nil {
		return err
	}
	return nil
}

// InstallPackage install a package with a package manager
func InstallPackage(name string, cmd *exec.Cmd) (bool, error) {
	log.Infoln("Trying to install the package: ", name)
	if ok, err := runCommand(cmd); !ok {
		log.WithError(err).Errorln("Could not install the package. ")

		return false, err
	}
	log.Infoln("Package installed: ", name)

	return true, nil
}

// UninstallPackage uninstall a package with a package manager
func UninstallPackage(name string, cmd *exec.Cmd) (bool, error) {
	log.Infoln("Uninstalling package: ", name)
	if ok, err := runCommand(cmd); !ok {
		log.Errorln("Could not uninstall the package ", err)

		return false, err
	}
	log.Errorln("Package uninstalled: ", name)

	return true, nil
}

// ListPackages dumps all the installed packages via a package manager
func ListPackages(cmd *exec.Cmd) (bool, error) {
	log.Infoln("Listing packages")
	if ok, err := runCommand(cmd); !ok {
		log.Errorln("Could not list packages: ", err)

		return false, err
	}

	return true, nil
}
