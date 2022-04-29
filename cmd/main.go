package main

import (
	"log"
	"webdiff/internal/download"
	"webdiff/internal/rest"
)

func main() {
	var downloads = download.StartQueue()

	err := rest.StartApi(downloads)
	if err != nil {
		log.Fatalln(err)
	}
}
