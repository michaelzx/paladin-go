package gormkit

import (
	"github.com/jinzhu/gorm"
)

type Params map[string]interface{}

func RawPlus(db *gorm.DB, sqlTpl string, params interface{}) *gorm.DB {
	resolver := new(SqlTplResolver)
	err := resolver.Resolve(sqlTpl, params)
	if err != nil {
		clone := db.New()
		_ = clone.AddError(err)
		return clone
	}
	return db.Raw(resolver.Sql, resolver.Values...)
}
