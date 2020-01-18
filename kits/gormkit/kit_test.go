package gormkit

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
)

type AttrVO struct {
	Id       int64
	SiteId   int64
	Title    string
	FormType int
	OptsTxt  string
	Enable   int
}

func TestRawPlus(t *testing.T) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2fShanghai", // &charset=utf8
		"root",
		"root",
		"127.0.0.1",
		3306,
		"zky",
	)
	db, err := gorm.Open("mysql", connStr)
	if err == nil {
		db.SingularTable(true)
	} else {
		panic(err)
	}
	AttrGetList := `select attr.id as id
,attr.form_type as form_type
,attr.title
,GROUP_CONCAT(ao.opt_txt,"、") as opts_txt
,if(sa.attr_id is null,0,1) as enable
from attr 
left join attr_opt as ao on ao.attr_id=attr.id
left join site_attr as sa on (sa.site_id=#{SiteId} and sa.attr_id=attr.id)
where attr.site_id in (0,#{SiteId})
group by attr.id`
	list := make([]AttrVO, 0, 0)
	result := RawPlus(db, AttrGetList, &struct {
		SiteId int64
	}{1})
	if result.Error != nil {
		if result.RecordNotFound() {
			return
		} else {
			panic(result.Error)
		}
	}
	result.Scan(&list)
	fmt.Printf("%#v", list)
}
