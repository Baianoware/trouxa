package parser

import (
	"github.com/Bainoware/trouxa/internal/entities"
	"strings"
)

func splitLines(data []byte) []string {
	return strings.Split(string(data), "\n")
}

// Parse the packages' list to entities.Packages
func Parse(data []byte) []entities.Package {
	lines := splitLines(data)
	var packages []entities.Package
	for _, line := range lines {
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		packages = append(packages, entities.Package{Name: line})
	}
	return packages
}
