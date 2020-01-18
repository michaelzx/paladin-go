package ginkit

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/michaelzx/paladin-go/errs"
)

func ParamToBool(c *gin.Context, paramName string) bool {
	return c.Param(paramName) == "true"
}
func ParamToInt64(c *gin.Context, paramName string) int64 {
	str := c.Param(paramName)
	if str == "" {
		panic(errs.ParamsNotExist.Suffix(paramName))
	}
	i64, err := strconv.ParseInt(c.Param(paramName), 10, 64)
	if err != nil {
		panic(errs.ParamsErr.Suffix(paramName + "，必须是数字"))
	}
	if i64 == 0 {
		panic(errs.ParamsErr.Suffix(paramName + "，不能为0"))
	}
	return i64
}
func ParamToInt(c *gin.Context, paramName string) int {
	str := c.Param(paramName)
	if str == "" {
		panic(errs.ParamsNotExist.Suffix(paramName))
	}
	i, err := strconv.Atoi(c.Param(paramName))
	if err != nil {
		panic(errs.ParamsErr.Suffix(paramName + "，必须是数字"))
	}
	if i == 0 {
		panic(errs.ParamsErr.Suffix(paramName + "，不能为0"))
	}
	return i
}

func GetPageSize(c *gin.Context) int {
	size := ParamToInt(c, "PageSize")
	if size <= 0 {
		size = 10
	}
	return size
}
func GetPageNum(c *gin.Context) int {
	num := ParamToInt(c, "PageNum")
	if num <= 0 {
		num = 1
	}
	return num
}
