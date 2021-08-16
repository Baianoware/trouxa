package manager

import "testing"

//TODO Remove the usage of vim as test package
func TestManagerPacman_InstallPackage(t *testing.T) {
	p := new(ManagerPacman)
	if _, ok := p.InstallPackage("vim"); !ok {
		t.Fail()
	}
}
func TestManagerPacman_UninstallPackage(t *testing.T) {
	p := new(ManagerPacman)
	if _, ok := p.UninstallPackage("vim"); !ok {
		t.Fail()
	}
}
