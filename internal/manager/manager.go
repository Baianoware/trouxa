package manager

import (
	"log"
	"os"
	"os/exec"
)

// Manager Interface with methods used to each package manager implementation
type Manager interface {
	InstallPackage(name string) bool
	UninstallPackage(name string) bool
}

// runCommand Execute a command and send the output to default Stdout and Stderr
func runCommand(cmd *exec.Cmd) (error, bool) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println("Trying to run the command")
	err := cmd.Run()
	if err != nil {
		log.Println("could not run the command ", err)
		return err, false
	}
	log.Println("Command executed")
	return nil, true
}
