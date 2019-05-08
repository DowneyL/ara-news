package user

import (
	"ara-news/components/mysql"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id       int    `orm:"column(id)" json:"id"`
	Name     string `orm:"column(name)" json:"name"`
	OpenId   string `orm:"column(open_id)" json:"open_id"`
	NickName string `orm:"column(nick_name)" json:"nick_name"`
	Type     int    `orm:"column(type)" json:"type"`
	IsRoot   int    `orm:"column(is_root)" json:"is_root"`
	CreateAt string `orm:"column(create_at)" json:"create_at"`
	UpdateAt string `orm:"column(update_at)" json:"update_at"`
}

func init() {
	orm.RegisterModel(new(Model))
}

func InitQuerySetter(genre ...string) orm.QuerySeter {
	var alias string
	if len(genre) > 0 && genre[0] != "" {
		alias = genre[0]
	}

	return mysql.GetQuerySetter(&Model{}, alias)
}

func (m *Model) TableName() string {
	return "user"
}

func FindAll() (users []Model, err error) {
	qs := InitQuerySetter()
	cols := mysql.GetCols(&Model{}, "IsRoot")
	_, err = qs.All(&users, cols...)
	return
}

func FindOneById(id int) (user Model, err error) {
	qs := InitQuerySetter()
	cols := mysql.GetCols(&Model{}, "IsRoot")
	err = qs.Filter("id", id).One(&user, cols...)
	return
}
