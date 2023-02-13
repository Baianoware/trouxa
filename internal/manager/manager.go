package manager

import (
	"bytes"
	"os"
	"os/exec"
)

func pipe(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}

func run(cmd *exec.Cmd) (bool, error) {
	write := func(b *bytes.Buffer) (int, error) {
		return os.Stdout.Write(b.Bytes())
	}

	mem := bytes.NewBuffer(nil)
	cmd.Stdin = os.Stdin
	cmd.Stdout = mem
	cmd.Stderr = mem

	if err := cmd.Run(); err != nil {
		if _, err := write(mem); err != nil {
			return false, err
		}

		return false, err
	}

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
	if ok, err := run(cmd); !ok || err != nil {
		return false, err
	}

	return true, nil
}

// UninstallPackage uninstall a package with a package manager
func UninstallPackage(name string, cmd *exec.Cmd) (bool, error) {
	if ok, err := run(cmd); !ok || err != nil {
		return false, err
	}

	return true, nil
}

// ListPackages lists all the installed packages via a package manager
func ListPackages(cmd *exec.Cmd) (bool, error) {
	pipe(cmd)
	if err := cmd.Run(); err != nil {
		return false, err
	}

	return true, nil
}
