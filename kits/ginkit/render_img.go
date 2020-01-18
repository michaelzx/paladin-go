package ginkit

import (
	"net/http"
)

type ImageRender struct {
	ImgBytes       []byte
	ImgContentType []string
}

func (r ImageRender) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	_, err := w.Write(r.ImgBytes)
	return err

}

var JpgContentType = []string{"image/jpeg"}

func (r ImageRender) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = r.ImgContentType
	}
}
