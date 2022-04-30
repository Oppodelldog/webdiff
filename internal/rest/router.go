package rest

import (
	"github.com/guiferpa/gody/v2"
	"github.com/guiferpa/gody/v2/rule"
	"github.com/julienschmidt/httprouter"
	"webdiff/internal/download"
)

func Router(router *httprouter.Router, downloads chan<- download.Request) {
	var (
		validator = mustCreateValidator()
	)

	router.GET("/rest/sessions", SessionHandler())
	router.GET("/rest/session/:session/urls", SessionUrlsHandler())
	router.GET("/rest/files/", AllFilesHandler())
	router.GET("/rest/files/:session", FilesBySessionHandler())
	router.GET("/rest/file/:session/:id", FileHandler())
	router.GET("/rest/diff/:sessionA/:idA/:sessionB/:idB", DiffHandler())
	router.POST("/rest/download", downloadHandler(validator, downloads))

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
