package parser

import (
	"github.com/Bainoware/trouxa/internal/entities"
	"log"
	"strings"
)

func splitLines(data []byte) []string {
	return strings.Split(string(data), "\n")
}

func Parse(data []byte) []entities.Package {
	lines := splitLines(data)
	var packages []entities.Package
	for _, line := range lines {
		if line == "" {
			log.Fatal("Error in the configuration file: a line is empty")
		}
		packages = append(packages, entities.Package{Name: line})
	}
	return packages
}
