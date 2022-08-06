package rest

import (
	"encoding/json"
	"fmt"
	"github.com/Oppodelldog/webdiff/internal/files"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func SessionHandler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		sessions, err := files.Sessions()
		if err != nil {
			http.Error(writer, fmt.Sprintf("cannot read sessions"), http.StatusInternalServerError)

			return
		}

		writer.Header().Set("Content-Type", "application/json")

		json.NewEncoder(writer).Encode(struct {
			Sessions []string `json:"sessions"`
		}{
			Sessions: sessions,
		})
	}
}

func SessionUrlsHandler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		session := params.ByName("session")
		urls, err := files.SessionUrls(session)
		if err != nil {
			http.Error(writer, fmt.Sprintf("cannot read session '%s' files", session), http.StatusInternalServerError)

			return
		}

		writer.Header().Set("Content-Type", "application/json")

		json.NewEncoder(writer).Encode(struct {
			Urls []string `json:"urls"`
		}{
			Urls: urls,
		})
	}
}
