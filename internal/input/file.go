package input

import (
	"io/ioutil"
)

func loadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func getPackagesFromFile(packagesPath string) ([]byte, error) {
	return loadFile(packagesPath)
}
