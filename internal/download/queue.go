package download

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/url"
	"time"
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
	var requests = make(chan Request, 1000)

	go func() {
		for request := range requests {
			var id = request.Id
			time.Sleep(1 * time.Second)
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
		}
	}()

	go func() {
		var queueSize int
		t := time.NewTicker(time.Millisecond * 500)
		for range t.C {
			var actualQueueSize = len(requests)
			if queueSize != actualQueueSize {
				client.NotifyQueueUpdate(actualQueueSize)
				queueSize = actualQueueSize
			}
		}
	}()

	return requests
}
