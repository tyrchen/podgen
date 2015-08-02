package commands

import (
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "host the generated site locally",
	Long:  `host the generated site locally so that you could look and feel it before pushing`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
