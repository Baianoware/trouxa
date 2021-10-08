package parser

import (
	"testing"

	"github.com/Bainoware/trouxa/internal/entities"
)

func TestSplitName(t *testing.T) {
	name := splitName("vim=1.0")
	if name != "vim" {
		t.Fail()
	}
}

func TestSplitVersion(t *testing.T) {
	version := splitVersion("vim=1.0")
	if version != "1.0" {
		t.Fail()
	}
}

func TestParser(t *testing.T) {
	data := Parse([]byte("vim=1.0\npython=3\nnano\nmysql="))
	expected := []entities.Package{
		{Name: "vim", Version: "1.0"},
		{Name: "python", Version: "3"},
		{Name: "nano"},
		{Name: "mysql"},
	}
	for i, e := range expected {
		if e.Name != data[i].Name || e.Version != data[i].Version {
			t.Fail()
		}
	}
}
