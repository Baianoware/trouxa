package commands

import (
	"os"

	"github.com/Bainoware/trouxa/internal/commander"
	"github.com/Bainoware/trouxa/internal/input"
	"github.com/Bainoware/trouxa/internal/manager"
	"github.com/Bainoware/trouxa/internal/parser"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var packageManager string
var pathFilePackages string

var RootCmd = &cobra.Command{
	Use:   "trouxa",
	Short: "Trouxa is a simple application to install and remove packages at once by just using a simple text file.",
	Long: `
Trouxa is a simple application to install and remove packages at once by just using a simple text file.
`,
	Version: "0.1.0",
	Run: func(command *cobra.Command, args []string) {
		installPackage := func(name string, commander commander.Commander) bool {
			ok, _ := manager.InstallPackage(name, commander.BuildInstallCommand(name))
			return ok
		}
		addTo := func(list *[]string, value string) {
			*list = append(*list, value)
		}
		var successInstalledPackages []string
		var failInstalledPackages []string
		commander := commander.FromName(packageManager)
		if commander == nil {
			log.Fatalln("This package manager is not supported or invalid for your system!")
		}
		if dump, _ := command.Flags().GetBool("dump"); dump {
			log.Debugln("Entered dump, will do nothing else")
			manager.ListPackages(commander.DumpPackages())
			os.Exit(0)
		}
		data, err := input.GetData(pathFilePackages)
		if err != nil {
			log.WithError(err).Fatalln() // TODO
		}
		packages := parser.Parse(data)
		for _, package_ := range packages {
			if !installPackage(package_.Name, commander) {
				log.Infoln("Fail to install: ", package_.Name)
				addTo(&failInstalledPackages, package_.Name)
				continue
			}
			log.Infoln("Successful installation of: ", package_.Name)
			addTo(&successInstalledPackages, package_.Name)
		}
		log.Infoln("Total success package's installation:", successInstalledPackages)
		log.Infoln("Total fail package's installation:", failInstalledPackages)
		os.Exit(0)
	},
}

func init() {
	RootCmd.Flags().StringVarP(
		&packageManager,
		"manager",
		"m",
		"",
		"Package manager to download your packages",
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
	RootCmd.Flags().Bool(
		"dump",
		false,
		"List all installed packages in this package manager and exit")
}
