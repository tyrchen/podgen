package commands

import (
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push the podcast site to github pages",
	Long:  `push the podcast site to github pages. Github will host the static files for you`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
