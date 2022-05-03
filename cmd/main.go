package main

import (
	"fmt"
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

	fmt.Printf("access webdiff webapp on %s", "http://localhost:12345/webapp")
	err := http.ListenAndServe(":12345", r)
	if err != nil {
		panic(err)
	}
}
