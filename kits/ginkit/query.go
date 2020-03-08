package ginkit

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/paladin-go/errs"
	"strconv"
)

func QueryToBool(c *gin.Context, paramName string) bool {
	return c.Query(paramName) == "true"
}
func QueryToInt64(c *gin.Context, paramName string) int64 {
	i64 := QueryToInt64WithZero(c, paramName)
	if i64 == 0 {
		panic(errs.QueryErr.Suffix(paramName + "，不能为0"))
	}

	return i64
}
func QueryToInt64WithZero(c *gin.Context, paramName string) int64 {
	str := c.Query(paramName)
	if str == "" {
		panic(errs.QueryNotExist.Suffix(paramName))
	}
	i64, err := strconv.ParseInt(c.Query(paramName), 10, 64)
	if err != nil {
		panic(errs.QueryErr.Suffix(paramName + "，必须是数字"))
	}
	return i64
}
func QueryToInt(c *gin.Context, paramName string) int {
	str := c.Query(paramName)
	if str == "" {
		panic(errs.QueryNotExist.Suffix(paramName))
	}
	i, err := strconv.Atoi(c.Query(paramName))
	if err != nil {
		panic(errs.QueryErr.Suffix(paramName + "，必须是数字"))
	}
	if i == 0 {
		panic(errs.QueryErr.Suffix(paramName + "，不能为0"))
	}
	return i
}
