package manager

import (
	"log"
	"os"
	"os/exec"
)

// Manager Interface with methods used to each package manager implementation
type Manager interface {
	InstallPackage(name string) (bool, error)
	UninstallPackage(name string) (bool, error)
}

// runCommand Execute a command and send the output to default Stdout and Stderr
func runCommand(cmd *exec.Cmd) (bool, error) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println("Trying to run the command")
	err := cmd.Run()
	if err != nil {
		log.Println("Could not run the command ", err)
		return false, err
	}
	log.Println("Command executed")
	return true, nil
}
