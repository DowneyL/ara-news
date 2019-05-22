package help_document_category

import (
	"ara-news/components/mysql"
	"ara-news/helper"
	helpValidator "ara-news/validators/help"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id             int64            `json:"id"`
	AttributeSetId helper.Platform  `json:"-"`
	Seq            int              `json:"seq"`
	IsHidden       int              `json:"-"`
	CreatedAt      helper.Timestamp `json:"-"`
	UpdatedAt      helper.Timestamp `json:"-"`
}

func init() {
	orm.RegisterModel(new(Model))
}

func (m *Model) TableName() string {
	return "help_document_category"
}

func InitQuerySetter(genre ...string) orm.QuerySeter {
	var alias string
	if len(genre) > 0 && genre[0] != "" {
		alias = genre[0]
	}

	return mysql.GetQuerySetter(&Model{}, alias)
}

func NewModel(category helpValidator.Category) Model {
	var model Model
	model.AttributeSetId = helper.GetAttrSetId(category.Platform)
	model.Seq = category.Seq
	model.CreatedAt = helper.Now()
	model.UpdatedAt = helper.Now()
	return model
}

func Insert(category helpValidator.Category) (int64, error) {
	model := NewModel(category)
	o := mysql.GetOrmer("master")

	return o.Insert(&model)
}

func TransactionInsert(o orm.Ormer, c helpValidator.Category) (int64, error) {
	model := NewModel(c)
	i, err := o.Insert(&model)
	if err != nil {
		_ = o.Rollback()
		return 0, err
	}

	return i, nil
}

func Exist(id int64) bool {
	qs := InitQuerySetter()

	return qs.Filter("id", id).Exist()
}
