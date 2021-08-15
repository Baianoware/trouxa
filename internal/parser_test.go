package internal

import "testing"

func TestParserConfigurationFile(t *testing.T) {
	parser := Parser{Sep: ":"}
	parser.ParsePackagesFile("../config/packages.txt")
}
