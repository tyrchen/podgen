package commands

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"

	"podgen/utils"
)

const (
	ITEM_TEMPLATE = `
- title: change_me
  link: chang_me.mp3
  image: change_me.png
  description: >
    change_me
  pubdate: %s
  guid: %s
`
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate a new item into items.yml",
	Long:  "Create a new episode from the template, appending it to items.yml. Please edit that file",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile("items.yml", os.O_APPEND|os.O_WRONLY, 0600)
		utils.CheckError(err)

		defer f.Close()

		guid := uuid.NewV4().String()
		pubdate := time.Now().Format(time.RFC3339)

		content := fmt.Sprintf(ITEM_TEMPLATE, pubdate, guid)
		_, err = f.WriteString(content)
		utils.CheckError(err)

		log.Println("Item info generated. Please modify items.yml.")
	},
}
