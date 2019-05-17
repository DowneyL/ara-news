package news_info_extend

import (
	"ara-news/boot"
	"ara-news/components/mysql"
	"ara-news/helper"
	newsValidator "ara-news/validators/news"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id        int64  `json:"id,omitempty"`
	Nid       int64  `json:"-"`
	ViewCount *int64 `json:"view_count,omitempty"`
}

func init() {
	orm.RegisterModel(new(Model))
}

func (m *Model) TableName() string {
	return "news_info_extend"
}

func InitQuerySetter(genre ...string) orm.QuerySeter {
	var alias string
	if len(genre) > 0 && genre[0] != "" {
		alias = genre[0]
	}

	return mysql.GetQuerySetter(&Model{}, alias)
}

func NewModel(nid int64) Model {
	var (
		model     Model
		viewCount int64 = 0
	)
	model.Nid = nid
	model.ViewCount = &viewCount

	return model
}

func FindByNid(nid int64, cols ...string) (Model, error) {
	var model Model
	qs := InitQuerySetter()
	err := qs.Filter("nid", nid).One(&model, cols...)

	return model, err
}

func TransactionInsert(o orm.Ormer, nid int64) error {
	model := NewModel(nid)
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
		_, err := qs.Filter("nid__in", query.Ids).All(&models, cols...)
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

func DeleteByNid(nid int64) (int64, error) {
	o := mysql.GetOrmer("master")

	return o.Delete(&Model{Nid: nid}, "nid")
}
