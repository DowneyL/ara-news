package news_info

import (
	"ara-news/components/mysql"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id             int64  `json:"id,omitempty"`
	Cid            int64  `json:"cid,omitempty"`
	AttributeSetId int    `json:"attribute_set_id,omitempty"`
	Seq            int    `json:"seq,omitempty"`
	IsHidden       int    `json:"is_hidden,omitempty"`
	Author         string `json:"author,omitempty"`
	CoverUrl       string `json:"cover_url,omitempty"`
	PublishedAt    int64  `json:"published_at,omitempty"`
	CreatedAt      int64  `json:"created_at,omitempty"`
	UpdatedAt      int64  `json:"updated_at,omitempty"`
}

func init() {
	orm.RegisterModel(new(Model))
}

func (m *Model) TableName() string {
	return "news_info"
}

func InitQuerySetter(genre ...string) orm.QuerySeter {
	var alias string
	if len(genre) > 0 && genre[0] != "" {
		alias = genre[0]
	}

	return mysql.GetQuerySetter(&Model{}, alias)
}
