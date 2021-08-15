package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trouxa pacman -p packages.txt",
	Short: "Trouxa is a simple application to install and remove a bunch of packages at once by just using a simple text file.",
	Long: `
Trouxa is a simple application to install and remove a bunch of packages at once by just using a simple text file.
`,
	Version: "0.0.1",
}
