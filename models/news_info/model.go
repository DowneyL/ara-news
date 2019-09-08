package news_info

import (
	"ara-news/boot"
	"ara-news/components/mysql"
	newsValidator "ara-news/components/validators/news"
	"ara-news/helper"
	"ara-news/models/news_category"
	"errors"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id             int64            `json:"id,omitempty"`
	Cid            int64            `json:"-"`
	AttributeSetId helper.Platform  `json:"-"`
	Seq            int              `json:"seq,omitempty"`
	IsHidden       int              `json:"is_hidden,omitempty"`
	Author         string           `json:"author,omitempty"`
	CoverUrl       string           `json:"cover_url,omitempty"`
	PublishedAt    helper.Timestamp `json:"-"`
	CreatedAt      helper.Timestamp `json:"-"`
	UpdatedAt      helper.Timestamp `json:"-"`
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

func NewModel(info newsValidator.Info) (Model, error) {
	var model Model
	category, err := news_category.FindByCode(info.CategoryCode)
	if err != nil {
		return model, err
	}
	if category.Id == 0 {
		return model, errors.New("category code not exist")
	}

	now := helper.NewTimestamp()
	model.Cid = category.Id
	model.AttributeSetId = helper.GetAttrSetId(info.Platform)
	model.Seq = info.Seq
	if info.IsHidden {
		model.IsHidden = 1
	}
	model.Author = info.Author
	model.CoverUrl = info.CoverUrl
	if info.PublishedAt != 0 {
		model.PublishedAt = helper.NewTimestamp(info.PublishedAt)
	} else {
		model.PublishedAt = now
	}
	model.CreatedAt = now
	model.UpdatedAt = now

	return model, nil
}

func Exist(id int64) bool {
	qs := InitQuerySetter()

	return qs.Filter("id", id).Exist()
}

func FindById(id int64, cols ...string) (Model, error) {
	var model Model
	qs := InitQuerySetter()
	err := qs.Filter("id", id).Filter("is_hidden", 0).One(&model, cols...)

	return model, err
}

func TransactionInsert(o orm.Ormer, info newsValidator.Info) (int64, error) {
	model, err := NewModel(info)
	if err != nil {
		return 0, err
	}
	nid, err := o.Insert(&model)
	if err != nil {
		_ = o.Rollback()
		return 0, err
	}

	return nid, nil
}

func FindLimit(query newsValidator.Query, cols ...string) ([]*Model, error) {
	var (
		models []*Model
		err    error
	)
	qs := InitQuerySetter()

	if len(query.Ids) > 0 {
		_, err = qs.Filter("id__in", query.Ids).Filter("is_hidden", 0).All(&models)
		return models, err
	}

	if query.Author != "" {
		qs = qs.Filter("author", query.Author)
	}
	if query.CategoryId != 0 {
		qs = qs.Filter("cid", query.CategoryId)
	}
	if query.Platform != "" {
		qs = qs.Filter("attribute_set_id", helper.GetAttrSetId(query.Platform))
	}

	qs = qs.Filter("is_hidden", 0).Filter("published_at__lte", helper.Now())

	if query.OrderBy != "" {
		orders, err := helper.GetOrmOrders(query.OrderBy)
		if err != nil {
			return models, err
		}
		qs = qs.OrderBy(orders...)
	} else {
		qs = qs.OrderBy("-published_at", "-id")
	}

	pagination := boot.GetPagination()
	i, err := qs.Limit(pagination.Size, pagination.Limit).All(&models, cols...)
	if i == 0 {
		err = orm.ErrNoRows
	}

	return models, err
}

func TransactionDeleteById(o orm.Ormer, id int64) (int64, error) {
	i, err := o.Delete(&Model{Id: id})
	if err != nil {
		_ = o.Rollback()
		return 0, err
	}

	return i, nil
}

func DeleteById(id int64) (int64, error) {
	o := mysql.GetOrmer("master")

	return o.Delete(&Model{Id: id})
}
