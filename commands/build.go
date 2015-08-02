package commands

import (
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build the podcast site",
	Long:  `build the podcast site, generate html files against template`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
