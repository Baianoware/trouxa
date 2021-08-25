package main

import (
	"fmt"
	commands "github.com/Bainoware/trouxa/cmd"
	"os"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
