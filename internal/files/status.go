package files

import (
	"net/http"
	"net/url"
)

type PageDownloadStatus struct {
	Token    string
	DateTime int64
	Download string
	Response PageDownloadResponse
	Request  PageDownloadRequest
}
type PageDownloadResponse struct {
	StatusCode       int
	Status           string
	Header           http.Header
	ContentLength    int64
	TransferEncoding []string
}
type PageDownloadRequest struct {
	Method string
	URL    *url.URL
	Header http.Header
}
