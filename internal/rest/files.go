package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"webdiff/internal/files"
)

func AllFilesHandler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fileList, err := files.Files("")
		if err != nil {
			http.Error(writer, fmt.Sprintf("cannot read all files"), http.StatusInternalServerError)

			return
		}

		writeFilesResponse(writer, fileList)
	}
}

func FilesBySessionHandler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var name = sanitizeFilename(params.ByName("session"))

		fileList, err := files.Files(name)
		if err != nil {
			http.Error(writer, fmt.Sprintf("cannot read files"), http.StatusInternalServerError)

			return
		}

		writeFilesResponse(writer, fileList)
	}
}

func writeFilesResponse(writer http.ResponseWriter, fileList files.FileEntries) {
	type FileEntries []files.FileEntry
	type FileEntry struct {
		File    string `json:"file"`
		Session string `json:"session"`
	}

	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(
		struct {
			Files FileEntries `json:"files"`
		}{
			Files: FileEntries(fileList),
		})
}
