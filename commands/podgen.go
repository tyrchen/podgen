package commands

import (
	"github.com/spf13/cobra"
)

const (
	DEFAULT_TMPL    = "https://github.com/tyrchen/podgen-basic"
	TEMPLATE_PATH   = "template"
	GH_PAGES        = "gh-pages"
	TARGET_PATH     = "build"
	ASSETS_PATH     = "assets"
	MAX_DESCRIPTION = 96
)

// Root command
var PodgenCmd = &cobra.Command{
	Use:   "podgen",
	Short: "podgen builds your podcast site",
	Long: `podgen is the command to build your awesome podcast site

podgen is a fast and flexible static site generator for podcast. If you'd like to publish a 
podcast in iTunes you need to host your media files yourself and provide a rss to iTunes. podgen
helps you to do it quite easy with just a few commands. You don't need a server to host the files,
podgen leverages the powerful github pages.

Steps to create a podcast site:

1. Create a public repo in github and clone it to a directory.
2. Init a podcast site by using "podgen init", under that directory.
3. Modify the "*.yml" files and copy the images/mp3 to desired sub directory, modify CNAME file for custom domain. See https://help.github.com/articles/setting-up-a-custom-domain-with-github-pages/.
4. Build the site by using "podgen build".
5. Look and feel the site by using "podgen server" (optional).
6. Push the site by using "podgen push".
7. You're all site. Now you can browse your site and register the rss in iTunes.

Next when you have new podcast you just modify "items.yml" and copy the related files. Then do 4-6.

Complete documentation is available at http://github.com/tyrchen/podgen.`,
}

func Execute() {
	PodgenCmd.AddCommand(initCmd, buildCmd, serverCmd, pushCmd)
	PodgenCmd.Execute()
}
