package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/yosssi/gohtml"
	"net/http"
	"strconv"
	"webdiff/internal/files"
)

func FileHandler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var (
			id         = sanitizeFilename(params.ByName("id"))
			session    = sanitizeFilename(params.ByName("session"))
			filterName = request.URL.Query().Get("filter")
			pretty     = request.URL.Query().Get("pretty")
		)

		archivedFile, err := files.FileFiltered(session, id, filterName)
		if err != nil {
			if isFilterErr(err) {
				json.NewEncoder(writer).Encode(files.ErrorResponse{Error: err.Error()})
				return
			} else {
				fmt.Println(err)
				http.Error(writer, fmt.Sprintf("cannot read file"), http.StatusInternalServerError)
			}

			return
		}

		if v, err := strconv.ParseBool(pretty); err == nil && v {
			archivedFile.Content = gohtml.FormatBytes(archivedFile.Content)
		}

		type ArchivedFile struct {
			Id      string `json:"name"`
			Session string `json:"session"`
			Content []byte `json:"content"`
		}

		json.NewEncoder(writer).Encode(ArchivedFile(*archivedFile))
	}
}

func isFilterErr(err error) bool {
	return errors.Is(err, files.ErrFilterNoMatch) ||
		errors.Is(err, files.ErrFilterInvalid) ||
		errors.Is(err, files.ErrParsingFailed) ||
		errors.Is(err, files.ErrUnknownFilterType)
}
