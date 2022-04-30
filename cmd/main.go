package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"webdiff/internal/download"
	"webdiff/internal/rest"
	"webdiff/internal/webapp"
)

func main() {
	var downloads = download.StartQueue()

	r := httprouter.New()
	rest.Router(r, downloads)
	webapp.Handler(r)

	err := http.ListenAndServe(":12345", r)
	if err != nil {
		panic(err)
	}
}
