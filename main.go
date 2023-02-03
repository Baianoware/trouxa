package main

import (
	"os"

	"github.com/Bainoware/trouxa/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
