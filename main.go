package main

import (
	"log"

	commands "github.com/Bainoware/trouxa/cmd"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
