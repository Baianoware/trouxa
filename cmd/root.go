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

func addPackagFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(
		&pathPackages,
		"path",
		"p",
		"./packages.txt",
		"The file used by Trouxa to download de packages to your system")
}
