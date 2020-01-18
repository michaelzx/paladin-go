package ginkit

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/paladin-go/errs"
	"github.com/michaelzx/paladin-go/validatorx"
)

func GetDTO(c *gin.Context, dtoPtr interface{}) {
	if bindErr := c.BindJSON(dtoPtr); bindErr != nil {
		if vErrs, ok := bindErr.(validatorx.ValidationErrors); ok {
			for _, vErr := range vErrs {
				e := vErr.(validatorx.FieldError)
				panic(errs.Common.Suffix(e.ErrMsg()))
			}
		} else {
			panic(errs.NewBadRequestErrorWithCode(10000, bindErr.Error()))
		}
	}
}
