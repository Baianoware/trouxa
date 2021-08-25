package commands

import (
	"github.com/spf13/cobra"
)

var pathPackages string

var RootCmd = &cobra.Command{
	Use:   "trouxa",
	Short: "Trouxa is a simple application to install and remove packages at once by just using a simple text file.",
	Long: `
Trouxa is a simple application to install and remove packages at once by just using a simple text file.
`,
	Version: "0.0.1",
}

// TODO Add a flag to trouxa root command to define packages.txt instead create a new flag for each command
//func init()  {
//	RootCmd.Flags().StringVarP(
//		&pathPackages,
//		"path",
//		"p",
//		"./packages.txt",
//		"File's path used by Trouxa to download the packages to your system")
//}
