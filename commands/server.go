package commands

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"gopkg.in/fsnotify.v1"

	"podgen/utils"
)

var (
	port int
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "host the generated site locally",
	Long:  `host the generated site locally so that you could look and feel it before pushing`,
	Run: func(cmd *cobra.Command, args []string) {

		addr := fmt.Sprintf(":%d", port)
		url := fmt.Sprintf("http://localhost:%d/index.html", port)

		go func() {
			time.Sleep(1)
			log.Printf("Starting your default browser with %s\n", url)
			open.Start(url)
		}()

		processSignal(watchFiles())

		log.Printf("Serving static content on %s\n", addr)
		http.ListenAndServe(addr, http.FileServer(http.Dir("./build")))

	},
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", DEFAULT_PORT, "port to listen")
}

func processSignal(watcher *fsnotify.Watcher) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Println("Shutting down...")
		watcher.Close()
		os.Exit(-1)
	}()

}

func watchFiles() (watcher *fsnotify.Watcher) {
	watcher, err := fsnotify.NewWatcher()
	utils.CheckError(err)
	changes := fsnotify.Write | fsnotify.Rename | fsnotify.Create

	needBuild := make(chan bool, 1)

	go func() {
		modified := false
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&changes != 0 {
					modified = true
				}
			case <-time.After(time.Second * 1):

				if modified {
					modified = false
					needBuild <- true
				}

			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	go func() {
		for {
			select {
			case build := <-needBuild:
				if build {
					log.Println("Rebuilding the project...")
					execute()
				}
			}
		}
	}()

	watchDirs := func(path string, info os.FileInfo, err error) error {
		stat, err := os.Stat(path)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			watcher.Add(path)
		}
		return nil
	}

	err = filepath.Walk("template", watchDirs)
	utils.CheckError(err)

	watcher.Add(".")
	utils.CheckError(err)

	return
}
