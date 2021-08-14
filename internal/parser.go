// Package internal is the code exclusive made for the
package internal

import (
	"github.com/Bainoware/trouxa/internal/entities"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func loadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func splitLines(data []byte) []string {
	return strings.Split(string(data), "\n")
}
func splitNameVersion(data []string) [][]string {
	var lines [][]string
	for _, line := range data {
		lines = append(lines, strings.Split(line, ":"))
	}
	return lines
}

func LoadConfigurationFile() []entities.Package {
	dataFromFile := loadFile("../config/packages.txt")
	split := splitNameVersion(splitLines(dataFromFile))
	var packages []entities.Package
	for _, line := range split {
		if line[0] == "" || line[1] == "" {
			log.Fatal("Error in the configuration file")
			os.Exit(1)
		}
		packages = append(packages, entities.Package{Name: line[0], Version: line[1]})
	}
	return packages
}
