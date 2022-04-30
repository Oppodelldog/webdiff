package rest

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/guiferpa/gody/v2"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"webdiff/internal/download"
)

type DownloadRequest struct {
	Id      string `json:"id"`
	Url     string `json:"url" validate:"not_empty"`
	Session string `json:"session" validate:"not_empty"`
}

func downloadHandler(validator *gody.Validator, downloads chan<- download.Request) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var downloadRequest DownloadRequest

		err := json.NewDecoder(request.Body).Decode(&downloadRequest)
		if err != nil {
			http.Error(writer, "error decoding Request", http.StatusBadRequest)

			return
		}
		defer request.Body.Close()

		if isValidated, err := validator.Validate(downloadRequest); err != nil {
			http.Error(writer, fmt.Sprintf("error in validation: %v", err), http.StatusBadRequest)

			return
		} else if !isValidated {
			http.Error(writer, "validation failed", http.StatusBadRequest)

			return
		}

		downloadToken := uuid.New().String()
		downloads <- download.Request{
			Token:   downloadToken,
			Id:      downloadRequest.Id,
			Url:     downloadRequest.Url,
			Session: downloadRequest.Session,
		}

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(struct {
			Token string `json:"token"`
		}{Token: downloadToken})
		if err != nil {
			log.Printf("could not encode response to client: %v", err)
		}
	}
}
