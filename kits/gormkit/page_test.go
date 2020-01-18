package gormkit

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
	"zky-server-app/ent"
)

func TestAbc(t *testing.T) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2fShanghai", // &charset=utf8
		"ddy_zx_io_cn",
		"CRsCRfcaytATHzYt",
		"116.62.218.204",
		3306,
		"ddy_zx_io_cn",
	)
	db, err := gorm.Open("mysql", connStr)
	if err == nil {
		db.SingularTable(true)
	} else {
		panic(err)
	}

	// service层参考 start
	page := &PageVO{
		PageNum:  3,
		PageSize: 7,
		List:     &[]ent.Goods{}, // 可根据返回字段，自定义pogo
	}
	err = page.Get(db, "select * from goods where shop_id = #{ShopId}", &Params{
		"ShopId": 1,
	})
	if err != nil {
		panic(err)
	}
	// service层参考 end

	jsonBytes, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}
