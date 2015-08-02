package utils

import (
	"log"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	log.Printf("Path %s exists but with error info: %s", err)
	return true
}
