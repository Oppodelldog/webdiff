package rest

import (
	"github.com/guiferpa/gody/v2"
	"github.com/guiferpa/gody/v2/rule"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"webdiff/internal/download"
)

func StartApi(downloads chan<- download.Request) error {

	return http.ListenAndServe(":12345", Router(downloads))
}

func Router(downloads chan<- download.Request) http.Handler {
	var (
		router    = httprouter.New()
		validator = mustCreateValidator()
	)

	router.GET("/rest/session", SessionHandler())
	router.GET("/rest/files/", AllFilesHandler())
	router.GET("/rest/files/:session", FilesBySessionHandler())
	router.GET("/rest/file/:session/:id", FileHandler())
	router.GET("/rest/diff/:sessionA/:idA/:sessionB/:idB", DiffHandler())
	router.POST("/rest/download", downloadHandler(validator, downloads))

	return router
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
