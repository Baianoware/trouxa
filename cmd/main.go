package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var pathPackages string

var rootCmd = &cobra.Command{
	Use:   "trouxa -p packages.txt",
	Short: "Trouxa is a simple application to install and remove a bunch of packages at once by just using a simple text file.",
	Long: `
Trouxa is a simple application to install and remove a bunch of packages at once by just using a simple text file.
`,
	Run: func(command *cobra.Command, args []string) {
		//TODO Implements here
		fmt.Println("The package path is", pathPackages)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	//TODO Handle error
	rootCmd.Flags().StringVarP(
		&pathPackages,
		"path",
		"p",
		"./packages.txt",
		"The file used by Trouxa to download de packages to your system")
	Execute()
}
