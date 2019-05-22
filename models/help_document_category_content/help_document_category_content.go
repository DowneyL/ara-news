package help_document_category_content

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
	LangType  lang.Types       `json:"-"`
	Text      string           `json:"text"`
	CreatedAt helper.Timestamp `json:"-"`
	UpdatedAt helper.Timestamp `json:"-"`
}

func init() {
	orm.RegisterModel(new(Model))
}

func (m *Model) TableName() string {
	return "help_document_category_content"
}

func InitQuerySetter(genre ...string) orm.QuerySeter {
	var alias string
	if len(genre) > 0 && genre[0] != "" {
		alias = genre[0]
	}

	return mysql.GetQuerySetter(&Model{}, alias)
}

func NewModel(cid int64, c helpValidator.CateContent) Model {
	return Model{
		Cid:       cid,
		LangType:  lang.GetLangId(c.Lang),
		Text:      c.Text,
		CreatedAt: helper.Now(),
		UpdatedAt: helper.Now(),
	}
}

func Insert(cid int64, cc helpValidator.CateContent) (int64, error) {
	o := mysql.GetOrmer("master")
	model := NewModel(cid, cc)

	return o.Insert(&model)
}

func TransactionInsert(o orm.Ormer, cid int64, c helpValidator.CateContent) (int64, error) {
	model := NewModel(cid, c)
	i, err := o.Insert(&model)
	if err != nil {
		_ = o.Rollback()
		return 0, err
	}

	return i, nil
}
