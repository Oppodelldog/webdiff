package webapp

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func Handler(r *httprouter.Router) {
	r.GET("/webapp", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		f, err := Filesystem("WEBDIFF_WEBAPP_DIR", templatesFS).Open("templates/index.html")
		if err != nil {
			http.Error(writer, "error", http.StatusInternalServerError)

			return
		}
		defer f.Close()
		content, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(writer, "error", http.StatusInternalServerError)

			return
		}

		writer.Write(content)
	})
	r.ServeFiles("/webapp/*filepath", http.FS(Filesystem("WEBDIFF_WEBAPP_DIR", assetsFS)))
}
