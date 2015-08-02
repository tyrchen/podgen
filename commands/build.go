package commands

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/andjosh/gopod"
	"github.com/spf13/cobra"

	"podgen/utils"
)

type Item struct {
	Title       string
	Description string
	Link        string
	PubDate     string
	Image       string
}

type Channel struct {
	Title       string
	Link        string
	Description string
	Image       string
	Copyright   string
	Language    string
	Author      string
	Categories  []string
	Page        int
	Twitter     string
	Linkedin    string
	Github      string
}

type PageTemplate struct {
	Info      Channel
	Home      string
	Current   Item
	Podcasts  []Item
	Paginator template.HTML
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build the podcast site",
	Long:  `build the podcast site, generate html files against template`,
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

// command implementation

func execute() {
	generatePages()
}

func generatePages() {
	channel := getChannelData("channel.yml")
	items := getItemData("items.yml")

	current := items[0]
	chopped_items := chopItems(items, channel.Page)
	len_chopped_items := len(chopped_items)

	generateRss(channel, items)

	pages := make([]int, len_chopped_items)
	for i := 1; i <= len_chopped_items; i++ {
		pages[i-1] = i
	}

	data, _ := ioutil.ReadFile(fmt.Sprintf("%s/index.tmpl", DEST_PATH))
	content := string(data[:])
	funcs := template.FuncMap{"alt": alt}
	t := template.Must(template.New("Podgen").Funcs(funcs).Parse(content))

	for i := 1; i <= len_chopped_items; i++ {
		var filename string
		var home string
		if i == 1 {
			filename = "index.html"
			home = "#current"
		} else {
			filename = fmt.Sprintf("page%d.html", i)
			home = "index.html"
		}
		f, err := os.Create(fmt.Sprintf("%s/%s", TARGET_PATH, filename))
		utils.CheckError(err)
		defer f.Close()

		err = t.Execute(f, PageTemplate{
			Info:      channel,
			Home:      home,
			Current:   current,
			Podcasts:  chopped_items[i-1],
			Paginator: generatePaginator(i, len_chopped_items),
		})
		utils.CheckError(err)
	}
}

func generateRss(channel Channel, items []Item) {
	c := gopod.ChannelFactory(channel.Title, channel.Link, channel.Description, channel.Image)
	c.SetiTunesExplicit("No")
	c.SetCopyright(channel.Copyright)
	c.SetiTunesAuthor(channel.Author)
	c.SetiTunesSummary(channel.Description)
	c.SetCategory(strings.Join(channel.Categories, ","))
	c.SetLanguage(channel.Language)

	for _, item := range items {
		url := path.Join(c.Link, "assets", item.Link)
		enclosure := gopod.Enclosure{
			Url:    url,
			Length: "0",
			Type:   "audio/mpeg",
		}
		c.AddItem(&gopod.Item{
			Title:         item.Title,
			Link:          url,
			Description:   item.Description,
			PubDate:       item.PubDate,
			TunesAuthor:   channel.Author,
			TunesSubtitle: item.Description,
			TunesSummary:  item.Description,
			Enclosure:     []*gopod.Enclosure{&enclosure},
		})
	}
	f, err := os.Create(fmt.Sprintf("%s/%s", TARGET_PATH, "rss.xml"))
	utils.CheckError(err)
	defer f.Close()
	f.Write(c.Publish())
}

func getChannelData(filename string) (channel Channel) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Cannot read file %s (%s)", filename, err)
	}
	err = yaml.Unmarshal(data, &channel)
	if err != nil {
		log.Fatalf("Cannot parse file %s (%s)", filename, err)
	}
	return
}

func getItemData(filename string) (items []Item) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Cannot read file %s (%s)", filename, err)
	}
	err = yaml.Unmarshal(data, &items)
	if err != nil {
		log.Fatalf("Cannot parse file %s (%s)", filename, err)
	}
	return
}

func chopItems(items []Item, page int) (chopped_items [][]Item) {

	length := len(items)
	j := 0
	for i := 1; i < length; i += page {
		chopped_items = append(chopped_items, items[i:int(math.Min(float64(i+page), float64(length)))])
		j += 1
	}
	return
}

func alt(x int) string {
	if x%2 == 0 {
		return "a"
	} else {
		return "b"
	}
}

// I cannot bear the golang template to do such a not-that-complicated paginator,
// thus I just do string concat myself...please tell me an elegant way to do so!!
func generatePaginator(curPage int, maxPage int) template.HTML {
	var data []string
	var pageName string
	var css_class string
	if curPage == 1 {
		data = append(data, "<li class='disabled'><a href='#'>&laquo;</a></li>")
	} else {
		if curPage-1 == 1 {
			pageName = "index.html"
		} else {
			pageName = fmt.Sprintf("page%d.html", (curPage - 1))
		}
		data = append(data, fmt.Sprintf("<li><a href='%s'>&laquo;</a></li>", pageName))
	}
	for i := 1; i <= maxPage; i++ {
		if i == curPage {
			css_class = "active"
		} else {
			css_class = ""
		}

		if i == 1 {
			pageName = "index.html"
		} else {
			pageName = fmt.Sprintf("page%d.html", i)
		}
		data = append(data, fmt.Sprintf("<li class='%s'><a href='%s'>%d</a></li>", css_class, pageName, i))
	}

	if curPage == maxPage {
		data = append(data, "<li class='disabled'><a href='#'>&raquo;</a></li>")
	} else {
		data = append(data, fmt.Sprintf("<li><a href='page%d.html'>&raquo;</a></li>", (curPage+1)))
	}

	return template.HTML(strings.Join(data, "\n"))
}
