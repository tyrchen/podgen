package commands

import (
	"github.com/codeskyblue/go-sh"
	"github.com/spf13/cobra"
)

const (
	DEFAULT_PUSH_MSG = "add new podcast."
)

var (
	push_msg string
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push the podcast site to github pages",
	Long:  `push the podcast site to github pages. Github will host the static files for you`,
	Run: func(cmd *cobra.Command, args []string) {
		session := sh.NewSession()
		gitCommit(session, push_msg, "master", true)
		session.SetDir("./build")
		gitCommit(session, push_msg, "gh-pages", true)
	},
}

func init() {
	pushCmd.Flags().StringVarP(&push_msg, "message", "m", DEFAULT_PUSH_MSG, "use the given message for git log")
}
