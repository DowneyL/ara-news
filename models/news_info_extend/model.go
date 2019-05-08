package news_info_extend

import (
	"ara-news/components/mysql"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id        int `orm:"column(id)"`
	Nid       int `orm:"column(nid)"`
	ViewCount int `orm:"column(view_count)"`
}

var qs orm.QuerySeter

func init() {
	orm.RegisterModel(new(Model))
	qs = mysql.GetQuerySetter(&Model{})
}

func (m *Model) TableName() string {
	return "news_info_extend"
}
