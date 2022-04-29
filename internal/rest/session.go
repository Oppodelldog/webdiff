package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"webdiff/internal/files"
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
