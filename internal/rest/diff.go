package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"path"
	"webdiff/internal/config"
	"webdiff/internal/diff"
)

func DiffHandler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var (
			sessionA  = sanitizeFilename(params.ByName("sessionA"))
			sessionB  = sanitizeFilename(params.ByName("sessionB"))
			fileA     = sanitizeFilename(params.ByName("fileA"))
			fileB     = sanitizeFilename(params.ByName("fileB"))
			filePathA = path.Join(config.DataDir(), sessionA, fileA)
			filePathB = path.Join(config.DataDir(), sessionB, fileB)
		)

		res, err := diff.Diff(filePathA, filePathB)
		if err != nil {
			http.Error(writer, fmt.Sprintf("cannot create diff"), http.StatusBadRequest)

			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			FileA    string `json:"file_a"`
			SessionA string `json:"session_a"`
			Diff     string `json:"diff"`
			FileB    string `json:"file_b"`
			SessionB string `json:"session_b"`
		}{
			FileA:    fileA,
			SessionA: sessionA,
			FileB:    fileB,
			SessionB: sessionB,
			Diff:     res,
		})

		if err != nil {
			log.Printf("error sending diff (%s,%s,%s,%s) response to client: %v", sessionA, fileA, sessionB, fileB, err)
		}
	}
}
