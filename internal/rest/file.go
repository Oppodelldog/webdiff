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
			file    = sanitizeFilename(params.ByName("file"))
			session = sanitizeFilename(params.ByName("session"))
		)

		archivedFile, err := files.File(session, file)
		if err != nil {
			http.Error(writer, fmt.Sprintf("cannot read files"), http.StatusInternalServerError)

			return
		}

		type ArchivedFile struct {
			Name    string `json:"name"`
			Session string `json:"session"`
			Content []byte `json:"content"`
		}

		json.NewEncoder(writer).Encode(ArchivedFile(*archivedFile))
	}
}
