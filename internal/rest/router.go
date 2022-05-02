package rest

import (
	"github.com/guiferpa/gody/v2"
	"github.com/guiferpa/gody/v2/rule"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"webdiff/internal/download"
	"webdiff/pkg/ws"
)

func Router(router *httprouter.Router, downloads chan<- download.Request, hub *ws.Hub) {
	var (
		validator = mustCreateValidator()
	)

	router.GET("/rest/sessions", SessionHandler())
	router.GET("/rest/session/:session/urls", SessionUrlsHandler())
	router.GET("/rest/files/", AllFilesHandler())
	router.GET("/rest/files/:session", FilesBySessionHandler())
	router.GET("/rest/file/:session/:id", FileHandler())
	router.GET("/rest/diff/:sessionA/:idA/:sessionB/:idB", DiffHandler())
	router.GET("/rest/filters", allFiltersHandler())
	router.POST("/rest/filter", upsertFilterHandler())
	router.DELETE("/rest/filter/:name", deleteFilterHandler())
	router.POST("/rest/download", downloadHandler(validator, downloads))

	router.GET("/ws", websocketHandler(hub))
}

func websocketHandler(hub *ws.Hub) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ws.ServeWs(hub, writer, request)
	}
}

func mustCreateValidator() *gody.Validator {
	var (
		validator = gody.NewValidator()
		err       = validator.AddRules(rule.NotEmpty, rule.Min)
	)

	if err != nil {
		panic(err)
	}

	return validator
}
