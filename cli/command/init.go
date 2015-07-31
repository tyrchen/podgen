package command

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	CHANNEL_FILE = "channel.yml"
	ITEMS_FILE   = "items.yml"
	GH_PAGES     = "gh_pages"
)

// InitCommand is a Command that initializes a new Vault server.
type InitCommand struct {
	Meta
}

func (c *InitCommand) Run(args []string) int {
	c.Ui.Output(fmt.Sprintf("Initial your podcast directory..."))

	if !c.exists("./.git") {
		c.Ui.Output(fmt.Sprintf("Please make sure you initialized" +
			" the git repository and pushed it to github."))
		os.Exit(-1)
	}

	if c.exists("channel.yml") {
		c.Ui.Output(fmt.Sprintf("Hmm..found existing 'channel.yml' - seems" +
			" you're on an already initialized directory."))
		os.Exit(-1)
	}
	c.Ui.Output(fmt.Sprintf("Done!"))

	return 0
}

func (c *InitCommand) Synopsis() string {
	return "Initialize a new podcast site at the current directory"
}

func (c *InitCommand) Help() string {
	helpText := `
Usage: podgen init [options]

  Initialize a new podcast site.

  This command create configuration files and directories for you to
  begin with.

  This command can't be called on an already-initialized podgen site.

`
	return strings.TrimSpace(helpText)
}

func (c *InitCommand) exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Printf("Path %s exists but with error info: %s", err)
	return true
}
