package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"webdiff/internal/files"
)

func FileHandler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var (
			id         = sanitizeFilename(params.ByName("id"))
			session    = sanitizeFilename(params.ByName("session"))
			filterName = request.URL.Query().Get("filter")
		)

		archivedFile, err := files.FileFiltered(session, id, filterName)
		if err != nil {
			http.Error(writer, fmt.Sprintf("cannot read file"), http.StatusInternalServerError)

			return
		}

		type ArchivedFile struct {
			Id      string `json:"name"`
			Session string `json:"session"`
			Content []byte `json:"content"`
		}

		json.NewEncoder(writer).Encode(ArchivedFile(*archivedFile))
	}
}
