package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"webdiff/internal/diff"
	"webdiff/internal/files"
)

func DiffHandler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var (
			sessionA  = sanitizeFilename(params.ByName("sessionA"))
			sessionB  = sanitizeFilename(params.ByName("sessionB"))
			idA       = sanitizeFilename(params.ByName("idA"))
			idB       = sanitizeFilename(params.ByName("idB"))
			filePathA = files.DownloadedFilePath(sessionA, idA)
			filePathB = files.DownloadedFilePath(sessionB, idB)
		)

		res, err := diff.Diff(filePathA, filePathB)
		if err != nil {
			http.Error(writer, fmt.Sprintf("cannot create diff"), http.StatusBadRequest)

			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			IdA      string `json:"id_a"`
			SessionA string `json:"session_a"`
			Diff     string `json:"diff"`
			IdB      string `json:"id_b"`
			SessionB string `json:"session_b"`
		}{
			IdA:      idA,
			SessionA: sessionA,
			IdB:      idB,
			SessionB: sessionB,
			Diff:     res,
		})

		if err != nil {
			log.Printf("error sending diff (%s,%s,%s,%s) response to client: %v", sessionA, idA, sessionB, idB, err)
		}
	}
}
