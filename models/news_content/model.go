package news_content

import (
	"ara-news/components/lang"
	"ara-news/components/mysql"
	newsValidator "ara-news/validators/news"
	"github.com/astaxie/beego/orm"
	"html/template"
)

type Model struct {
	Id      int64  `json:"id,omitempty"`
	Nid     int64  `json:"nid,omitempty"`
	Lang    int    `json:"lang,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
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
	model.Lang = lang.GetLangId(content.Lang)
	model.Title = template.HTMLEscapeString(content.Title)
	model.Content = template.HTMLEscapeString(content.Content)

	return model
}

func FindByNId(nid int64, cols ...string) (Model, error) {
	var model Model
	qs := InitQuerySetter()
	err := qs.Filter("nid", nid).One(&model, cols...)

	return model, err
}
