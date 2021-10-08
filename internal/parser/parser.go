package parser

import (
	"strings"

	"github.com/Bainoware/trouxa/internal/entities"
)

func splitLines(data []byte) []string {
	return strings.Split(string(data), "\n")
}

func splitName(line string) string {
	return strings.Split(line, "=")[0]
}

func splitVersion(line string) string {
	return strings.Split(line, "=")[1]
}

// Parse the packages' list to entities.Packages
func Parse(data []byte) []entities.Package {
	addPackageNameVersion := func(packages *[]entities.Package, line string) {
		*packages = append(*packages, entities.Package{Name: splitName(line), Version: splitVersion(line)})
	}
	addPackageName := func(packages *[]entities.Package, line string) {
		*packages = append(*packages, entities.Package{Name: splitName(line)})
	}
	lines := splitLines(data)
	var packages []entities.Package
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.ContainsRune(line, '=') && splitVersion(line) != "" {
			addPackageNameVersion(&packages, line)
		} else {
			addPackageName(&packages, line)
		}
	}
	return packages
}
