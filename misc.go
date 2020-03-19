package main

import (
	"log"
	"os"
)

// https://siongui.github.io/2017/03/28/go-create-directory-if-not-exist/
// Create a directory if it does not exist. Otherwise do nothing.
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}
