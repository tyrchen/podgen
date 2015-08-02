package utils

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
