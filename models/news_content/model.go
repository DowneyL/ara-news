package news_content

import (
	"ara-news/boot"
	"ara-news/components/lang"
	"ara-news/components/mysql"
	newsValidator "ara-news/components/validators/news"
	"ara-news/helper"
	"github.com/astaxie/beego/orm"
	"html/template"
)

type Model struct {
	Id        int64      `json:"id,omitempty"`
	Nid       int64      `json:"-"`
	IsDefault *int       `json:"is_default,omitempty"`
	LangType  lang.Types `json:"-"`
	Title     string     `json:"title,omitempty"`
	Content   string     `json:"content,omitempty"`
}

func init() {
	orm.RegisterModel(new(Model))
}

func (m *Model) TableName() string {
	return "news_content"
}

func InitQuerySetter(genre ...string) orm.QuerySeter {
	var alias string
	if len(genre) > 0 && genre[0] != "" {
		alias = genre[0]
	}

	return mysql.GetQuerySetter(&Model{}, alias)
}

func NewModel(nid int64, content newsValidator.Content) Model {
	var model Model
	model.Nid = nid
	model.LangType = lang.GetLangId(content.Lang)
	model.Title = template.HTMLEscapeString(content.Title)
	model.Content = template.HTMLEscapeString(content.Content)
	var is = 0
	if content.Default {
		is = 1
	}
	model.IsDefault = &is

	return model
}

func FindAllByNId(nid int64, cols ...string) ([]*Model, error) {
	var models []*Model
	qs := InitQuerySetter()
	_, err := qs.Filter("nid", nid).All(&models, cols...)

	return models, err
}

func FindByNId(nid int64, cols ...string) (Model, error) {
	var model Model
	qs := InitQuerySetter()
	err := qs.Filter("nid", nid).Filter("lang_type", lang.GetLangId(boot.GetLang())).One(&model, cols...)

	return model, err
}

func TransactionInsert(o orm.Ormer, nid int64, content newsValidator.Content) error {
	model := NewModel(nid, content)
	_, err := o.Insert(&model)
	if err != nil {
		_ = o.Rollback()
		return err
	}

	return nil
}

func FindLimit(query newsValidator.Query, cols ...string) ([]*Model, error) {
	var (
		models []*Model
		err    error
	)
	qs := InitQuerySetter()

	if len(query.Ids) > 0 {
		_, err = qs.Filter("nid__in", query.Ids).Filter("lang_type", lang.GetLangId(boot.GetLang())).All(&models, cols...)
		return models, err
	}

	if query.OrderBy != "" {
		orders, err := helper.GetOrmOrders(query.OrderBy)
		if err != nil {
			return models, nil
		}
		qs.OrderBy(orders...)
	}
	pagination := boot.GetPagination()
	_, err = qs.Limit(pagination.Size, pagination.Limit).All(&models, cols...)

	return models, err
}

func Insert(nid int64, content newsValidator.Content) (int64, error) {
	o := mysql.GetOrmer("master")
	model := NewModel(nid, content)

	return o.Insert(&model)
}

func TransactionDeleteByNid(o orm.Ormer, nid int64) (int64, error) {
	i, err := o.Delete(&Model{Nid: nid}, "nid")
	if err != nil {
		_ = o.Rollback()
		return 0, err
	}

	return i, nil
}

func DeleteByNId(nid int64) (int64, error) {
	o := mysql.GetOrmer("master")

	return o.Delete(&Model{Nid: nid}, "nid")
}
