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
func runCommand(cmd *exec.Cmd) bool {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println("Trying to run the command")
	err := cmd.Run()
	if err != nil {
		log.Fatal("Could not run the command", err)
		return false
	}
	log.Println("Command executed")
	return true
}
