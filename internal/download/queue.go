package download

import (
	"log"
	"path"
	"webdiff/internal/config"
)

type Request struct {
	Token   string
	Id      string
	Url     string
	Session string
}

func StartQueue() chan<- Request {
	var requests = make(chan Request)
	var dataDir = config.DataDir()

	go func() {
		for request := range requests {
			go func(request Request) {
				var filename = path.Join(dataDir, request.Session, request.Id)
				err := Page(request.Url, filename)
				if err != nil {
					log.Printf("error downloading url='%s', session='%s': %v", request.Url, filename, err)
				}
			}(request)
		}
	}()

	return requests
}
