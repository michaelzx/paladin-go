package gormkit

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	"github.com/michaelzx/paladin-go/utils"
	"strings"
	"time"
)

func NewPageVO2(db *gorm.DB, mapList interface{}, resolver *SqlTplResolver, pageNum, pageSize int32) (*PageVO, error) {
	// return &PageVO{PageNum: pageNum, PageSize: pageSize, List: mapList}
	p := &PageVO{
		PageNum:  pageNum,
		PageSize: pageSize,
		List:     mapList,
	}
	err := p.Get4Map(db, resolver)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *PageVO) Get4Map(db *gorm.DB, resolver *SqlTplResolver) error {
	sqlBlocks := strings.Split(resolver.Sql, "from")
	countSql := "select count(1) from" + sqlBlocks[len(sqlBlocks)-1]
	countSql = strings.Replace(countSql, "\n", " ", -1)
	result := db.Raw(countSql, resolver.Values...).Count(&p.Total)
	if result.Error != nil {
		return result.Error
	}
	p.PageTotal = p.Total / p.PageSize
	if p.Total%p.PageSize > 0 {
		p.PageTotal = p.PageTotal + 1
	}
	if p.PageNum > p.PageTotal {
		p.PageNum = p.PageTotal
	}
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	if p.PageTotal == 0 {

		p.IsFirstPage = true
		p.IsLastPage = true
	} else {
		switch p.PageNum {
		case 1:
			p.IsFirstPage = true
		case p.PageTotal:
			p.IsLastPage = true
		}
	}
	skipRow := p.PageSize * (p.PageNum - 1)
	pageSql := fmt.Sprintf(`%s limit %d,%d`, resolver.Sql, skipRow, p.PageSize)
	pageSql = strings.Replace(pageSql, "\n", " ", -1)
	result = db.Raw(pageSql, resolver.Values...)
	if result.Error != nil {
		if result.RecordNotFound() {
			return nil
		} else {
			panic(result.Error)
		}
	}
	if !utils.IsPtr(p.List) {
		return errors.New("字段 List 必须是指针类型")
	}
	rows, err := result.Rows()
	if err != nil {
		panic(err)
	}
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	length := len(columns)
	list := make([]map[string]interface{}, 0)
	for rows.Next() {
		valMap := createBlankMapRow(length)
		if err := rows.Scan(valMap...); err != nil {
			panic(err)
		}
		row := make(map[string]interface{})
		fillMapRow(row, valMap, columns)
		list = append(list, row)
	}
	p.List = list
	// fmt.Printf("p.Total->%#v\n", p.Total)
	// fmt.Printf("p.PageNum->%#v\n", p.PageNum)
	// fmt.Printf("p.PageSize->%#v\n", p.PageSize)
	// fmt.Printf("p.PageTotal->%#v\n", p.PageTotal)
	// fmt.Printf("p.IsFirstPage->%#v\n", p.IsFirstPage)
	// fmt.Printf("p.IsLastPage->%#v\n", p.IsLastPage)
	// fmt.Printf("countSql->%#v\n", countSql)
	// fmt.Printf("pageSql->%#v\n", pageSql)
	return nil
}
func createBlankMapRow(length int) []interface{} {
	result := make([]interface{}, 0, length)
	for i := 0; i < length; i++ {
		var current interface{}
		current = struct{}{}
		result = append(result, &current)
	}
	return result
}
func fillMapRow(row map[string]interface{}, valMap []interface{}, columns []string) {
	for i := 0; i < len(columns); i++ {
		key := columns[i]
		val := *(valMap[i]).(*interface{})

		rowKey := strcase.ToCamel(key)
		if val == nil {
			row[rowKey] = nil
			continue
		}

		switch val.(type) {
		case int64:
			row[rowKey] = val.(int64)
		case string:
			row[rowKey] = val.(string)
		case time.Time:
			row[rowKey] = val.(time.Time)
		case []uint8:
			// fmt.Printf(" %s ->'%T' |%s\n", key, val, string(val.([]uint8)))
			if strings.HasPrefix(key, "attr_") {
				thisStr := string(val.([]uint8))
				var v interface{}
				err := json.Unmarshal([]byte(thisStr), &v)
				if err != nil {
					row[rowKey] = "error!!!"
				} else {
					row[rowKey] = v
				}
			} else {
				row[rowKey] = string(val.([]uint8))
			}
		default:
			// fmt.Printf("unsupport data type %s ->'%v' now\n", key, t)
		}
	}
}
