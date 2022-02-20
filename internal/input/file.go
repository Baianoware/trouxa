package input

import (
	"io/ioutil"
)

func getPackagesFromFile(packagesPath string) ([]byte, error) {
	return ioutil.ReadFile(packagesPath)
}
