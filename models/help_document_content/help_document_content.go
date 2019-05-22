package help_document_content

import (
	"ara-news/components/lang"
	"ara-news/components/mysql"
	"github.com/astaxie/beego/orm"
	"time"
)

type Model struct {
	Id        int64      `json:"id"`
	Cid       int64      `json:"cid"`
	LangType  lang.Types `json:"lang_type"`
	Question  string     `json:"question"`
	Answer    string     `json:"answer"`
	Seq       int        `json:"seq"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
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
