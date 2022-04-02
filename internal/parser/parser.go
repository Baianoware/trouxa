package parser

import (
	"log"
	"regexp"
	"strings"

	"github.com/Bainoware/trouxa/internal/entities"
)

// Parse the packages' list from config file to entities.Packages.
func Parse(data []byte) []entities.Package {
	// Gets all lines.
	lines := strings.Split(string(data), "\n")
	rule := regexp.MustCompile("^([a-zA-Z])*=?([0-9.])*$")

	packages := make([]entities.Package, 0)
	for index, line := range lines {
		// Valid if a line has the right format.
		// package=version
		if !rule.MatchString(line) {
			log.Fatalf("line %d, %s, has a invalid format", index, line)
		}

		// if package contains version indicator, put the version on package's structure.
		if strings.ContainsRune(line, '=') {
			name := strings.TrimSpace(strings.Split(lines[index], "=")[0])
			version := strings.TrimSpace(strings.Split(lines[index], "=")[1])

			packages = append(packages, entities.Package{Name: name, Version: version})
		} else {
			name := strings.TrimSpace(strings.Split(lines[index], "=")[0])

			packages = append(packages, entities.Package{Name: name})
		}
	}

	return packages
}
