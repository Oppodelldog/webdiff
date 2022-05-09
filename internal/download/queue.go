package download

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/url"
	"webdiff/internal/client"
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

				requestUrl, err := url.Parse(request.Url)
				if err != nil {
					var msg = fmt.Sprintf("error downloading url='%s': %v", request.Url, err)
					log.Print(msg)
					client.Log(msg, "ERROR")

					return
				}

				if len(id) > 4 && id[:4] == "gen:" {
					switch id {
					case "gen:hash_url":
						id = fmt.Sprintf("%x", md5.Sum([]byte(requestUrl.String())))
					case "gen:hash_path":
						id = fmt.Sprintf("%x", md5.Sum([]byte(requestUrl.Path)))
					}
				}

				var filename = files.DownloadedFilePath(request.Session, id)
				var statusFile = files.StatusFilePath(request.Session, id)

				err = Page(request.Token, requestUrl.String(), filename, statusFile)
				if err != nil {
					var msg = fmt.Sprintf("error downloading url='%s', session='%s': %v", request.Url, filename, err)
					log.Print(msg)
					client.Log(msg, "ERROR")
				}
			}(request)
		}
	}()

	return requests
}
