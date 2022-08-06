package rest

import (
	"encoding/json"
	"fmt"
	"github.com/Oppodelldog/webdiff/internal/files"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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
			log.Println(err)
			http.Error(writer, fmt.Sprintf("cannot read files"), http.StatusInternalServerError)

			return
		}

		writeFilesResponse(writer, fileList)
	}
}

type FileEntry struct {
	File    string `json:"file"`
	Session string `json:"session"`
}
type FileEntries []FileEntry

func writeFilesResponse(writer http.ResponseWriter, fileList files.FileEntries) {

	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(
		struct {
			Files FileEntries `json:"files"`
		}{
			Files: adoptList(fileList),
		})
}

func adoptList(list files.FileEntries) FileEntries {
	var resList FileEntries
	for _, entry := range list {
		resList = append(resList, FileEntry(entry))
	}

	return resList
}
