package command

import (
	"fmt"
	"strings"
)

// InitCommand is a Command that initializes a new Vault server.
type InitCommand struct {
	Meta
}

func (c *InitCommand) Run(args []string) int {
	c.Ui.Output(fmt.Sprintf("Initial your podcast directory"))

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
