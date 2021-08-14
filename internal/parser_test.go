package internal

import "testing"

func TestParserConfigurationFile(t *testing.T) {
	parser := Parser{Sep: ":"}
	parser.Parse()
}
