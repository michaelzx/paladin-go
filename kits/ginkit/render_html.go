package ginkit

import (
	"net/http"
)

type HtmlStringRender struct {
	HtmlString string
}

func (r HtmlStringRender) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	_, err := w.Write([]byte(r.HtmlString))
	return err

}

var htmlContentType = []string{"text/html; charset=utf-8"}

func (r HtmlStringRender) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = htmlContentType
	}
}
