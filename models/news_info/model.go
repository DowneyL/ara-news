package news_info

import (
	"ara-news/components/mysql"
	"ara-news/helper"
	"ara-news/models/news_category"
	newsValidator "ara-news/validators/news"
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
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

var model Model

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

func NewModel(info newsValidator.Info) (Model, error) {
	category, err := news_category.FindByCode(info.CategoryCode)
	if err != nil {
		return model, err
	}
	if category.Id == 0 {
		return model, errors.New("category code not exist")
	}

	now := time.Now().Unix()
	model.Cid = category.Id
	model.AttributeSetId = helper.GetAttrSetId(info.Platform)
	model.Seq = info.Seq
	if info.IsHidden {
		model.IsHidden = 1
	}
	model.Author = info.Author
	model.CoverUrl = info.CoverUrl
	if info.PublishedAt != 0 {
		model.PublishedAt = info.PublishedAt
	} else {
		model.PublishedAt = now
	}
	model.CreatedAt = now
	model.UpdatedAt = now

	return model, nil
}

func FindById(id int64, cols ...string) (Model, error) {
	qs := InitQuerySetter()
	err := qs.Filter("id", id).One(&model, cols...)

	return model, err
}
