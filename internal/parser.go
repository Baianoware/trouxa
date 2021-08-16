// Package internal is the code exclusive made for the
package internal

import (
	"github.com/Bainoware/trouxa/internal/entities"
	"io/ioutil"
	"log"
	"strings"
)

func loadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("could not find the packages file", err)
	}
	return data
}
func splitLines(data []byte) []string {
	return strings.Split(string(data), "\n")
}

//TODO Improve on unused struct
type Parser struct {
}

// ParsePackagesFile load the packages files
func (p *Parser) ParsePackagesFile(packagesFilePath string) []entities.Package {
	dataFromFile := loadFile(packagesFilePath)
	lines := splitLines(dataFromFile)
	var packages []entities.Package
	for _, line := range lines {
		if line == "" {
			log.Fatal("Error in the configuration file: a line is empty")
		}
		packages = append(packages, entities.Package{Name: line})
	}
	return packages
}
