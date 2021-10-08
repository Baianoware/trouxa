package main

import (
	"log"

	"github.com/Bainoware/trouxa/internal/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
