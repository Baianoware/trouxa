package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Bainoware/trouxa/internal/commander"
	"github.com/Bainoware/trouxa/internal/input"
	"github.com/Bainoware/trouxa/internal/manager"
	"github.com/Bainoware/trouxa/internal/parser"
	"github.com/spf13/cobra"
)

var (
	packageManager   string
	pathFilePackages string
)

func Info(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, msg+"\n", args...)
}

func Error(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

func Success(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, msg+"\n", args...)
	os.Exit(0)
}

func Fatal(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

var RootCmd = &cobra.Command{
	Use:   "trouxa",
	Short: "Trouxa is a simple application to install and remove packages at once by just using a simple text file.",
	Long: `
Trouxa is a simple application to install and remove packages at once by just using a simple text file.
`,
	Version: "0.1.0",
	Run: func(command *cobra.Command, args []string) {
		install := func(name string, commander commander.Commander) (bool, error) {
			ok, err := manager.InstallPackage(name, commander.BuildInstallCommand(name))
			if err != nil {
				return false, err
			}

			return ok, nil
		}
		addTo := func(list *[]string, value string) {
			*list = append(*list, value)
		}
		var installed []string
		var failed []string
		commander, err := commander.FromName(packageManager)
		if err != nil || commander == nil {
			Fatal("This package manager is not supported or invalid for your system!")
		}
		if list, _ := command.Flags().GetBool("list"); list {
			ok, err := manager.ListPackages(commander.ListPackages())
			if !ok {
				Fatal("Failed to list packages: %s", err)
			}
			Success("Listed packages successfully!")
		}
		data, err := input.GetData(pathFilePackages)
		if err != nil {
			Fatal("Failed to get data from file: %s", err)
		}
		packages := parser.Parse(data)
		for _, package_ := range packages {
			Info("Trying to install %s", package_.Name)
			if ok, err := install(package_.Name, commander); !ok || err != nil {
				Error("Fail to install %s", package_.Name)
				addTo(&failed, package_.Name)
				continue
			}
			Info("Successful installation of %s", package_.Name)
			addTo(&installed, package_.Name)
		}
		Info("Installation finished!")
		Info("Total success package's installation: %s", strings.Join(installed, ", "))
		Info("Total fail package's installation: %s", strings.Join(failed, ", "))
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
		Fatal("Could not bing 'manager' flag as required")
	}
	RootCmd.Flags().StringVarP(
		&pathFilePackages,
		"path",
		"p",
		"./packages.txt",
		"The file used by Trouxa to download de packages to your system")
	RootCmd.Flags().Bool(
		"list",
		false,
		"List all installed packages in this package manager and exit")
}
