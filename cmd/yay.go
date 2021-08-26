package commands

import (
	"log"
	"os"

	"github.com/Bainoware/trouxa/internal"
	"github.com/Bainoware/trouxa/internal/manager"
	"github.com/spf13/cobra"
)

var yayCmd = &cobra.Command{
	Example: "pacman -p packages.txt",
	Use:     "yay",
	Short:   "Use yay to download the packages",
	Long:    "Use yay for download the packages",
	Run: func(command *cobra.Command, args []string) {
		parser := new(internal.Parser)
		packages := parser.ParsePackagesFile(pathPackages)
		manager := new(manager.Yay)
		for _, package_ := range packages {
			if err, ok := manager.InstallPackage(package_.Name); !ok {
				log.Fatalln("Could not install ", package_.Name, ". aborting cuz' ", err)
			}
		}
		log.Println("All packages have installed!")
		os.Exit(0)
	},
}

func init() {
	yayCmd.Flags().StringVarP(
		&pathPackages,
		"path",
		"p",
		"./packages.txt",
		"The file used by Trouxa to download de packages to your system")
	RootCmd.AddCommand(yayCmd)
}
