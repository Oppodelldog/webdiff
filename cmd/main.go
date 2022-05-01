package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"webdiff/internal/client"
	"webdiff/internal/download"
	"webdiff/internal/rest"
	"webdiff/internal/webapp"
)

func main() {
	var r = httprouter.New()
	var downloads = download.StartQueue()
	var hub = client.StartWebsocketHub()

	rest.Router(r, downloads, hub)
	webapp.Handler(r)

	err := http.ListenAndServe(":12345", r)
	if err != nil {
		panic(err)
	}
}
