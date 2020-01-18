package gormkit

import (
	"fmt"
	"testing"
)

const tt = `select attr.id as id
,attr.form_type as form_type
,attr.title
,GROUP_CONCAT(ao.opt_txt,"、") as opts_txt
,if(sa.attr_id is null,0,1) as enable
from attr 
left join attr_opt as ao on ao.attr_id=attr.id
left join site_attr as sa on (sa.site_id=#{SiteId} and sa.attr_id=attr.id)
where attr.site_id in (0,#{SiteId})
group by attr.id`

type QueryStruct struct {
	SiteId  int64
	SiteId2 int64
}

func TestTpl(t *testing.T) {
	tpl := new(sqlTplResolver)
	err := tpl.Resolve(tt, &Params{
		"SiteId":  1,
		"SiteId2": 2,
	})
	if err == nil {
		fmt.Printf("%#v\n", tpl)
	} else {
		fmt.Println(err)
	}
	tpl = new(sqlTplResolver)
	err = tpl.Resolve(tt, &QueryStruct{
		SiteId:  1,
		SiteId2: 2,
	})
	if err == nil {
		fmt.Printf("%#v\n", tpl)
	} else {
		fmt.Println(err)
	}
}
