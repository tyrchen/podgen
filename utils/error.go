package utils

import (
	"log"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}
