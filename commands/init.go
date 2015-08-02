package commands

import (
	// "bytes"
	"fmt"
	// "log"
	"os"
	// "path/filepath"
	"strings"
	// "time"

	"github.com/codeskyblue/go-sh"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-gitconfig"

	"podgen/utils"
)

const (
	DEFAULT_TMPL = "https://github.com/tyrchen/podgen-basic"
	DEST_PATH    = "template"
	GH_PAGES     = "gh-pages"
)

var (
	template       string
	originUrl      string
	FILES_TO_CHECK = []string{"channel.yml", "items.yml", "build", "music", "template"}
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new podcast site on current directory",
	Long: `Initialize a new podcast site on current directiory.
Configuration files and site `,
	Run: func(cmd *cobra.Command, args []string) {
		originUrl = getOriginUrl()
		fmt.Printf("Current repo is %s, You're using template: %s\n", originUrl, template)
		execute()
	},
}

func init() {
	initCmd.Flags().StringVarP(&template, "template", "t", DEFAULT_TMPL, "Content type to create")
}

// command implementation

func execute() {
	if !utils.Exists("./.git") {
		fmt.Printf("'.git' is not found. Please create an empty github repo, clone it to you local directory and then run this command under the directory.")
		os.Exit(-1)
	}

	for _, filename := range FILES_TO_CHECK {
		if utils.Exists(filename) {
			fmt.Printf("Hmm...found existing '%s' - seems you're on an already initialized podcast directory. I cannot init it again.", filename)
			os.Exit(-1)
		}
	}
	getTemplate()
	createGhPages()

	fmt.Println("\nCongratulations, your podcast site is ready to use. Please modify the *.yml files and try to 'podgen build' your site!\n")
}

func getTemplate() {
	session := sh.NewSession()
	session.Command("git", "clone", "--depth=1", template, DEST_PATH).Run()
	removeFiles(session, ".git")
	mvFiles(session, "channel.yml", "items.yml", "music", ".gitignore")
	gitCommit(session, "initial podcast site", "master", true)
}

func createGhPages() {
	session := sh.NewSession()
	session.Command("git", "checkout", "--orphan", GH_PAGES).Run()
	session.Command("git", "rm", "-rf", ".").Run()
	session.Command("touch", "index.html").Run()

	gitCommit(session, "initial podcast site", "gh-pages", true)

	session.Command("git", "checkout", "master").Run()

	session.Command("git", "clone", "-b", GH_PAGES, originUrl, "build").Run()
}

func removeFiles(session *sh.Session, files ...string) {
	for _, filename := range files {
		session.Command("rm", "-rf", fmt.Sprintf("%s/%s", DEST_PATH, filename)).Run()
	}
}

func mvFiles(session *sh.Session, files ...string) {
	for _, filename := range files {
		session.Command("mv", fmt.Sprintf("%s/%s", DEST_PATH, filename), ".").Run()
	}
}

func gitCommit(session *sh.Session, message string, branch string, setUpstream bool) {
	session.Command("git", "add", ".").Run()
	session.Command("git", "commit", "-a", "-m", message).Run()
	if setUpstream {
		session.Command("git", "push", "-u", "origin", branch).Run()
	} else {
		session.Command("git", "push", "origin", branch).Run()
	}

}

func getOriginUrl() string {
	originUrl, err := gitconfig.OriginURL()
	utils.CheckError(err)
	if !strings.HasPrefix(originUrl, "git@") {
		fmt.Printf("Please clone the repo with SSL clone URL. Otherwise the repo cannot be modified (Origin url is %s in .git/config)\n", originUrl)
		os.Exit(-1)
	}
	return originUrl
}
