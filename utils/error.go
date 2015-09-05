package utils

import (
	"log"
	"os"
)

func CheckError(err error, messages ...string) {
	if err != nil {
		log.Println(err)
		if len(messages) > 0 {
			log.Println(messages)
		}

		os.Exit(-1)
	}
}
