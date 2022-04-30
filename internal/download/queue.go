package download

import (
	"crypto/md5"
	"fmt"
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
				var id = request.Id
				if len(id) == 0 {
					id = fmt.Sprintf("%x", md5.Sum([]byte(request.Url)))
				}
				var filename = files.DownloadedFilePath(request.Session, id)
				var statusFile = files.StatusFilePath(request.Session, id)
				err := Page(request.Token, request.Url, filename, statusFile)
				if err != nil {
					log.Printf("error downloading url='%s', session='%s': %v", request.Url, filename, err)
				}
			}(request)
		}
	}()

	return requests
}
