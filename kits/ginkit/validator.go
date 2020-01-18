package ginkit

import (
	"fmt"
	"reflect"
	"regexp"
	"sync"
	"github.com/michaelzx/paladin-go/validatorx"
)

// func init() {
// 	binding.Validator = &DefaultValidator{}
// }

type DefaultValidator struct {
	once     sync.Once
	validate *validatorx.Validate
}

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
func (v *DefaultValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

// Engine returns the underlying validator engine which powers the default
// Validator instance. This is useful if you want to register custom validations
// or struct level validations. See validator GoDoc for more info -
// https://godoc.org/gopkg.in/go-playground/validatorx.v8
func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validatorx.New()
		_ = v.validate.RegisterValidation("mobile", func(fl validatorx.FieldLevel) bool {
			v := fl.Field().String()
			reg := `^1\d{10}$`
			rgx := regexp.MustCompile(reg)
			return rgx.MatchString(v)
		})
		_ = v.validate.RegisterValidation("int", func(fl validatorx.FieldLevel) bool {
			p := fl.Param()
			v := fl.Field().String()
			reg := `^\d+$`
			if p != "" {
				reg = fmt.Sprintf(`^\d{%s}$`, p)
			}
			rgx := regexp.MustCompile(reg)
			return rgx.MatchString(v)
		})
		v.validate.SetTagName("valid")
	})
}
