package news_info

import (
	"ara-news/components/mysql"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id             int    `orm:"column(id)"`
	Cid            int    `orm:"column(cid)"`
	AttributeSetId int    `orm:"column(attribute_set_id)"`
	Seq            int    `orm:"column(seq)"`
	IsHidden       int    `orm:"column(is_hidden)"`
	Author         string `orm:"column(author)"`
	CoverUrl       string `orm:"column(cover_url)"`
	PublishedAt    int    `orm:"column(published_at)"`
	CreatedAt      int    `orm:"column(created_at)"`
	UpdatedAt      int    `orm:"column(updated_at)"`
}

var qs orm.QuerySeter

func init() {
	orm.RegisterModel(new(Model))
	qs = mysql.GetQuerySetter(&Model{})
}

func (m *Model) TableName() string {
	return "news_info"
}
