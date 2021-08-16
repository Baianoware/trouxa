package internal

import "testing"

func TestParserConfigurationFile(t *testing.T) {
	parser := new(Parser)
	parser.ParsePackagesFile("../config/packages.txt")
}
