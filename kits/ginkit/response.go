package ginkit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RJson(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, v)
}
func RXmlString(c *gin.Context, xml string) {
	c.Data(http.StatusOK, "application/xml; charset=utf-8", []byte(xml))
}
func RString(c *gin.Context, s string) {
	c.String(http.StatusOK, s)
}
