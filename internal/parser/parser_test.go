package parser

import (
	"testing"

	"github.com/Bainoware/trouxa/internal/entities"
)

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
