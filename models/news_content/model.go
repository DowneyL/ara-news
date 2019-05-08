package news_content

import (
	"ara-news/components/mysql"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id      int    `orm:"column(id)"`
	Nid     int    `orm:"column(nid)"`
	Lang    int    `orm:"column(lang)"`
	Title   string `orm:"column(title)"`
	Content string `orm:"column(content)"`
}

var qs orm.QuerySeter

func init() {
	orm.RegisterModel(new(Model))
	qs = mysql.GetQuerySetter(&Model{})
}

func (m *Model) TableName() string {
	return "news_content"
}
