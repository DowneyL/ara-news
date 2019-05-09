package news_category

import (
	"ara-news/boot"
	"ara-news/components/mysql"
	"ara-news/validators"
	"github.com/astaxie/beego/orm"
	"time"
)

//type Model struct {
//	Id int64 `orm:"column(id)"`
//	Code string `orm:"column(code)"`
//	Seq int `orm:"column(seq)"`
//	Icon string `orm:"column(icon)"`
//	NameZh string `orm:"column(name_zh)"`
//	NameEn string `orm:"column(name_en)"`
//	CreatedAt int64 `orm:"column(created_at)"`
//	UpdatedAt int64 `orm:"column(updated_at)"`
//}

type Model struct {
	Id        int64  `json:"id,omitempty"`
	Code      string `json:"code,omitempty"`
	Seq       int    `json:"seq,omitempty"`
	Icon      string `json:"icon,omitempty"`
	NameZh    string `json:"name_zh,omitempty"`
	NameEn    string `json:"name_en,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func init() {
	orm.RegisterModel(new(Model))
}

func (m *Model) TableName() string {
	return "news_category"
}

func InitQuerySetter(genre ...string) orm.QuerySeter {
	var alias string
	if len(genre) > 0 && genre[0] != "" {
		alias = genre[0]
	}

	return mysql.GetQuerySetter(&Model{}, alias)
}

func FindAll() []*Model {
	qs := InitQuerySetter()
	var categories []*Model
	pagination := boot.GetPagination()
	i, _ := qs.OrderBy("-created_at").Limit(pagination.Size, pagination.Limit).All(&categories)
	if i == 0 {
		return nil
	}

	return categories
}

func Insert(category validators.NewsCategory) (int64, error) {
	var model Model
	model.Seq = category.Seq
	model.Code = category.Code
	model.Icon = category.Icon
	model.NameZh = category.NameZH
	model.NameEn = category.NameEN
	now := time.Now().Unix()
	model.CreatedAt = now
	model.UpdatedAt = now
	o := mysql.GetOrmer("master")

	return o.Insert(&model)
}

func DeleteById(id int64) (int64, error) {
	o := mysql.GetOrmer("master")

	return o.Delete(&Model{Id: id})
}

func DeleteByIds(cids validators.CategoryIds) (int64, error) {
	ids := cids.Ids
	qs := InitQuerySetter("master")

	return qs.Filter("id__in", ids).Delete()
}
