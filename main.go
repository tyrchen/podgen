package main

import (
	// "fmt"
	"log"
	"os"
	// "time"

	"io/ioutil"

	// "github.com/andjosh/gopod"
	"gopkg.in/yaml.v2"

	"github.com/tyrchen/podgen/cli"
)

type Item struct {
	Title       string
	Description string
	Link        string
	PubDate     string
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
}

func GetChannelData(filename string) (channel Channel) {
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

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
