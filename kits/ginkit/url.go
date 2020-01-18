package ginkit

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

func UrlFull(c *gin.Context) string {
	scheme := "http://"
	if c.Request.TLS != nil {
		scheme = "https://"
	}
	return scheme + c.Request.Host + c.Request.RequestURI
}

func UrlCutQuery(fullUrl string, querys ...string) string {
	u, _ := url.Parse(fullUrl)
	if u.RawQuery == "" {
		return fullUrl
	}
	oldQueryStr := "?" + u.RawQuery
	vs := u.Query()
	for _, q := range querys {
		delete(vs, q)
	}
	newQueryArr := make([]string, 0, len(vs))
	for k, v := range vs {
		item := k + "=" + strings.Join(v, ",")
		newQueryArr = append(newQueryArr, item)
	}
	newQueryStr := ""
	if len(newQueryArr) > 0 {
		newQueryStr = "?" + strings.Join(newQueryArr, "&")
	}
	return strings.Replace(fullUrl, oldQueryStr, newQueryStr, 1)
}
