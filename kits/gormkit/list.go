package gormkit

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
	"github.com/michaelzx/paladin-go/utils"
)

type ListItem interface {
}

func NewListVO(db *gorm.DB, list interface{}, sqlTpl string, params interface{}) error {
	if !utils.IsPtr(list) {
		return errors.New("字段 List 必须是指针类型")
	}
	resolver := new(sqlTplResolver)
	err := resolver.Resolve(sqlTpl, params)
	if err != nil {
		return err
	}
	resolver.Sql = strings.Replace(resolver.Sql, "\n", " ", -1)
	result := db.Raw(resolver.Sql, resolver.Values...)
	switch {
	case result.RecordNotFound():
		return nil
	case result.Error != nil:
		return result.Error
	default:
		result.Scan(list)
	}
	return nil
}
