package commands

import (
	"log"
	"os"

	"github.com/Bainoware/trouxa/internal"
	"github.com/Bainoware/trouxa/internal/manager"
	"github.com/spf13/cobra"
)

var pacmanCmd = &cobra.Command{
	Example: "pacman -p packages.txt",
	Use:     "pacman",
	Short:   "Use pacman to download the packages",
	Long:    "Use pacman for download the packages",
	Run: func(command *cobra.Command, args []string) {
		parser := new(internal.Parser)
		packages := parser.ParsePackagesFile(pathPackages)
		manager := new(manager.Pacman)
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
	addPackagFlag(pacmanCmd)
	RootCmd.AddCommand(pacmanCmd)
}
