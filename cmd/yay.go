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
			if ok, err := manager.InstallPackage(package_.Name); !ok {
				log.Fatalln("Could not install ", package_.Name, ". aborting cuz' ", err)
			}
		}
		log.Println("All packages have installed!")
		os.Exit(0)
	},
}

func init() {
	addPackagFlag(yayCmd)
	RootCmd.AddCommand(yayCmd)
}
