package utils

import (
	"log"
	"os"
	"strings"
)

func CheckError(err error, messages ...string) {
	if err != nil {
		log.Println(err)
		if len(messages) > 0 {
			log.Println(strings.Join(messages, " "))
		}

		os.Exit(-1)
	}
}
