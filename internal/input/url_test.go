package input

import (
	"testing"
)

func TestURLPackages(t *testing.T) {
	data, _ := getPackagesFromURL("https://pastebin.com/raw/ysHUVswx")
	if len(data) == 0 {
		t.Fail()
	}
}
