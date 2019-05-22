package help_document_content

import (
	"ara-news/components/lang"
	"ara-news/components/mysql"
	"ara-news/helper"
	helpValidator "ara-news/validators/help"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id        int64            `json:"id"`
	Cid       int64            `json:"cid"`
	LangType  lang.Types       `json:"lang_type"`
	Question  string           `json:"question"`
	Answer    string           `json:"answer"`
	Seq       int              `json:"seq"`
	CreatedAt helper.Timestamp `json:"-"`
	UpdatedAt helper.Timestamp `json:"-"`
}

func init() {
	orm.RegisterModel(new(Model))
}

func (m *Model) TableName() string {
	return "help_document_content"
}

func InitQuerySetter(genre ...string) orm.QuerySeter {
	var alias string
	if len(genre) > 0 && genre[0] != "" {
		alias = genre[0]
	}

	return mysql.GetQuerySetter(&Model{}, alias)
}

func NewModel(content helpValidator.Content) Model {
	return Model{
		Cid:       content.Cid,
		LangType:  lang.GetLangId(content.Lang),
		Question:  content.Question,
		Answer:    content.Answer,
		Seq:       content.Seq,
		CreatedAt: helper.Now(),
		UpdatedAt: helper.Now(),
	}
}

func Insert(content helpValidator.Content) (int64, error) {
	model := NewModel(content)
	o := mysql.GetOrmer("master")

	return o.Insert(&model)
}
