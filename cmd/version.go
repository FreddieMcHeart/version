package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Version",
	Long:  `All software has versions. This is Version's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version Generator v.alpha.0.0 -- HEAD")
	},
}
