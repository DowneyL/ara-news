package news_category

import (
	"ara-news/boot"
	"ara-news/components/mysql"
	newsValidator "ara-news/components/validators/news"
	"ara-news/helper"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id        int64            `json:"id,omitempty"`
	Code      string           `json:"code,omitempty"`
	Seq       int              `json:"seq,omitempty"`
	Icon      string           `json:"icon,omitempty"`
	NameZh    string           `json:"-"`
	NameEn    string           `json:"-"`
	CreatedAt helper.Timestamp `json:"-"`
	UpdatedAt helper.Timestamp `json:"-"`
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

func FindLimit(query newsValidator.QueryCategory) ([]*Model, error) {
	var (
		categories []*Model
		err        error
	)
	qs := InitQuerySetter()

	if len(query.CIds) > 0 {
		_, err = qs.Filter("id__in", query.CIds).All(&categories)
		return categories, err
	}

	pagination := boot.GetPagination()
	if query.Code != "" {
		qs = qs.Filter("code", query.Code)
	}
	if query.NameEN != "" {
		qs = qs.Filter("name_en", query.NameEN)
	}
	if query.NameZH != "" {
		qs = qs.Filter("name_zh", query.NameZH)
	}

	i, err := qs.OrderBy("-seq", "-created_at").Limit(pagination.Size, pagination.Limit).All(&categories)
	if err != nil || i == 0 {
		return nil, err
	}

	return categories, nil
}

func FindById(id int64, cols ...string) (Model, error) {
	var model Model
	qs := InitQuerySetter()
	err := qs.Filter("id", id).One(&model, cols...)

	return model, err
}

func FindByCode(code string, cols ...string) (Model, error) {
	var model Model
	qs := InitQuerySetter()
	err := qs.Filter("code", code).One(&model, cols...)

	return model, err
}

func Insert(category newsValidator.Category) (int64, error) {
	var model Model
	model.Seq = category.Seq
	model.Code = category.Code
	model.Icon = category.Icon
	model.NameZh = category.NameZH
	model.NameEn = category.NameEN
	now := helper.NewTimestamp()
	model.CreatedAt = now
	model.UpdatedAt = now
	o := mysql.GetOrmer("master")

	return o.Insert(&model)
}

func UpdateById(id int64, category newsValidator.Category) (int64, error) {
	qs := InitQuerySetter("master")
	now := helper.NewTimestamp()
	return qs.Filter("id", id).Update(orm.Params{
		"code":       category.Code,
		"seq":        category.Seq,
		"icon":       category.Icon,
		"name_zh":    category.NameZH,
		"name_en":    category.NameEN,
		"updated_at": now,
	})
}

func UpdateNameEnById(id int64, categoryName newsValidator.UpdateNameEn) (int64, error) {
	qs := InitQuerySetter("master")
	now := helper.NewTimestamp()
	return qs.Filter("id", id).Update(orm.Params{
		"name_en":    categoryName.NameEN,
		"updated_at": now,
	})
}

func DeleteById(id int64) (int64, error) {
	o := mysql.GetOrmer("master")

	return o.Delete(&Model{Id: id})
}

func DeleteByIds(cIds newsValidator.CategoryIds) (int64, error) {
	ids := cIds.Ids
	qs := InitQuerySetter("master")

	return qs.Filter("id__in", ids).Delete()
}
