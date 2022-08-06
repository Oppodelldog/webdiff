package main

import (
	"fmt"
	"github.com/Oppodelldog/webdiff/internal/client"
	"github.com/Oppodelldog/webdiff/internal/download"
	"github.com/Oppodelldog/webdiff/internal/rest"
	"github.com/Oppodelldog/webdiff/internal/webapp"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	var r = httprouter.New()
	var hub = client.StartWebsocketHub()
	var downloads = download.StartQueue()

	rest.Router(r, downloads, hub)
	webapp.Handler(r)

	fmt.Printf("access webdiff webapp on %s", "http://localhost:12345/webapp")
	err := http.ListenAndServe(":12345", r)
	if err != nil {
		panic(err)
	}
}
