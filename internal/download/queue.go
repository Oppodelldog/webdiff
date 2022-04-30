package download

import (
	"log"
	"webdiff/internal/files"
)

type Request struct {
	Token   string
	Id      string
	Url     string
	Session string
}

func StartQueue() chan<- Request {
	var requests = make(chan Request)

	go func() {
		for request := range requests {
			go func(request Request) {
				var filename = files.DownloadedFilePath(request.Session, request.Id)
				var statusFile = files.StatusFilePath(request.Session, request.Id)
				err := Page(request.Token, request.Url, filename, statusFile)
				if err != nil {
					log.Printf("error downloading url='%s', session='%s': %v", request.Url, filename, err)
				}
			}(request)
		}
	}()

	return requests
}
