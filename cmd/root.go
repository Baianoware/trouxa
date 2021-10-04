package commands

import (
	"github.com/Bainoware/trouxa/internal/commander"
	"github.com/Bainoware/trouxa/internal/input"
	"github.com/Bainoware/trouxa/internal/manager"
	parser "github.com/Bainoware/trouxa/internal/parser"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var packageManager string
var pathFilePackages string

var RootCmd = &cobra.Command{
	Use:   "trouxa",
	Short: "Trouxa is a simple application to install and remove packages at once by just using a simple text file.",
	Long: `
Trouxa is a simple application to install and remove packages at once by just using a simple text file.
`,
	Version: "0.0.1",
	Run: func(command *cobra.Command, args []string) {
		commander := commander.FromName(packageManager)
		if commander == nil {
			log.Fatalln("This package manager is not supported!")
		}
		data, err := input.GetData(pathFilePackages)
		if err != nil {
			log.WithError(err).Fatalln() // TODO
		}
		packages := parser.Parse(data)
		for _, package_ := range packages {
			if ok, err := manager.InstallPackage(package_.Name, commander.BuildInstallCommand(package_.Name)); !ok {
				log.Fatalln("Could not install ", package_.Name, ". Aborting cuz' ", err)
			}
		}
		log.Infoln("All packages have installed!")
		os.Exit(0)
	},
}

func init() {
	RootCmd.Flags().StringVarP(
		&packageManager,
		"manager",
		"m",
		"",
		"Package manager to download your packages. apt | pacman | yay | apk",
	)
	err := RootCmd.MarkFlagRequired("manager")
	if err != nil {
		log.Fatalln("Could not bing 'manager' flag as required")
	}
	RootCmd.Flags().StringVarP(
		&pathFilePackages,
		"path",
		"p",
		"./packages.txt",
		"The file used by Trouxa to download de packages to your system")
}
