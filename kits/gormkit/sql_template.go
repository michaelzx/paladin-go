package gormkit

import (
	"bytes"
	"errors"
	"reflect"
	"regexp"
	"strings"
	"text/template"
	"github.com/michaelzx/paladin-go/utils"
)

type SqlTplResolver struct {
	Sql    string
	Values []interface{}
}

func (s *SqlTplResolver) Resolve(tplStr string, queryParams interface{}) error {
	if !utils.IsPtr(queryParams) {
		return errors.New("params 必须是指针类型")
	}
	t := template.New("sql")
	var err error
	t, err = t.Parse(tplStr)
	if err != nil {
		return err
	}
	var sql bytes.Buffer
	if err = t.Execute(&sql, queryParams); err != nil {
		return err
	}
	sqlStr := sql.String()
	// 正则找出，模板分析出中所有的 #{xxxx}
	tplParamsRegexp := regexp.MustCompile(`#{(?P<param>\w*)}`)
	tplParams := tplParamsRegexp.FindAllStringSubmatch(sqlStr, -1)
	// 校验 tplParams 是否 都在params中
	// 仅支持 Params 及 struct
	if p, ok := queryParams.(*Params); ok {
		s.Values = make([]interface{}, 0, len(tplParams))
		for _, tplParam := range tplParams {
			full := tplParam[0]
			short := tplParam[1]
			v, exist := (*p)[short]
			if !exist {
				panic(full + "：不在queryParams中")
			}
			sqlStr = strings.Replace(sqlStr, full, "?", 1)
			s.Values = append(s.Values, v)
		}
	} else if utils.IsStruct(queryParams) {
		rt := reflect.TypeOf(queryParams)
		rve := reflect.ValueOf(queryParams).Elem()
		fim := utils.StructFieldIdxMap(rt)
		for _, tplParam := range tplParams {
			full := tplParam[0]
			short := tplParam[1]
			fi, exist := fim[short]
			if !exist {
				panic(full + "：不在queryParams中")
			}
			v := rve.Field(fi).Interface()
			sqlStr = strings.Replace(sqlStr, full, "?", 1)
			s.Values = append(s.Values, v)
		}
	} else {
		return errors.New("queryParams 仅支持 Params 及 struct")
	}

	s.Sql = sqlStr
	return nil
}
