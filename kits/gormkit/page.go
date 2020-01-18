package gormkit

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"github.com/michaelzx/paladin-go/utils"
)

type PageCommonParams struct {
	PageNum  int32 `valid:"required~页码不能为空"`
	PageSize int32 `valid:"required~页数不能为空"`
}
type PageVO struct {
	PageNum     int32 // 第几页
	PageSize    int32 // 每页几条
	PageTotal   int32 // 总共几页
	Total       int32 // 总共几条
	IsFirstPage bool  // 是否是第一页
	IsLastPage  bool  // 是否是最后一页
	List        interface{}
}

func NewPageVO(db *gorm.DB, list interface{}, sqlTpl string, pageNum int32, pageSize int32, params interface{}) (*PageVO, error) {
	// return &PageVO{PageNum: pageNum, PageSize: pageSize, List: list}
	p := &PageVO{
		PageNum:  pageNum,
		PageSize: pageSize,
		List:     list,
	}
	err := p.Get(db, sqlTpl, params)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *PageVO) Get(db *gorm.DB, sqlTpl string, params interface{}) error {
	resolver := new(sqlTplResolver)
	// 先用text/template对sql解析一波
	err := resolver.Resolve(sqlTpl, params)
	if err != nil {
		return err
	}
	countSql := fmt.Sprintf(`select count(*) from (%s) as t`, resolver.Sql)
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
	result.Scan(p.List)
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
