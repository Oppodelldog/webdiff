package rest

import (
	"encoding/json"
	"fmt"
	"github.com/Oppodelldog/webdiff/internal/files"
	"github.com/google/go-cmp/cmp"
	"github.com/julienschmidt/httprouter"
	"github.com/yosssi/gohtml"
	"log"
	"net/http"
	"strconv"
)

type DiffResult struct {
	IdA      string `json:"id_a"`
	SessionA string `json:"session_a"`
	Diff     string `json:"diff"`
	IdB      string `json:"id_b"`
	SessionB string `json:"session_b"`
}

func DiffHandler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var (
			sessionA = sanitizeFilename(params.ByName("sessionA"))
			sessionB = sanitizeFilename(params.ByName("sessionB"))
			idA      = sanitizeFilename(params.ByName("idA"))
			idB      = sanitizeFilename(params.ByName("idB"))
			filter   = request.URL.Query().Get("filter")
			pretty   = request.URL.Query().Get("pretty")
		)

		contentA, errA := files.FileFiltered(sessionA, idA, filter)
		contentB, errB := files.FileFiltered(sessionB, idB, filter)
		if errA != nil || errB != nil {
			var respErrA, respErrB error
			if isFilterErr(errA) {
				respErrA = errA
			}
			if isFilterErr(errB) {
				respErrB = errB
			}

			if respErrA != nil || respErrB != nil {
				err := json.NewEncoder(writer).Encode(DiffResult{
					IdA:      idA,
					SessionA: sessionA,
					IdB:      idB,
					SessionB: sessionB,
					Diff:     fmt.Sprintf("error file A: %s\nerror file B: %s", respErrA, respErrB),
				})

				if err != nil {
					log.Printf("error sending diff (%s,%s,%s,%s) response to client: %v", sessionA, idA, sessionB, idB, err)
				}

				return
			}

			err := fmt.Errorf("error while loading files: Error File1: %v - Error File2: %v", errA, errB)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if v, err := strconv.ParseBool(pretty); err == nil && v {
			contentA.Content = gohtml.FormatBytes(contentA.Content)
			contentB.Content = gohtml.FormatBytes(contentB.Content)
		}

		res := cmp.Diff(string(contentA.Content), string(contentB.Content))
		err := json.NewEncoder(writer).Encode(DiffResult{
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
