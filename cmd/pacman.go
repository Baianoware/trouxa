package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var pathPackages string

var pacmanCmd = &cobra.Command{
	Use:   "pacman -p packages.txt",
	Short: "Use pacman to download the packages",
	Long:  "Use pacman for download the packages",
	Run: func(command *cobra.Command, args []string) {
		//TODO Implements here
		fmt.Println("The package path is", pathPackages)
	},
}

func init() {
	pacmanCmd.Flags().StringVarP(
		&pathPackages,
		"path",
		"p",
		"./packages.txt",
		"The file used by Trouxa to download de packages to your system")
	rootCmd.AddCommand(pacmanCmd)
}
