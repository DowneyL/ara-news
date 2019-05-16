package news

import (
	"ara-news/components/lang"
	"ara-news/models/news_category"
	newsValidator "ara-news/validators/news"
)

type CategoryList []*Category

type Category struct {
	news_category.Model
	Name        map[string]string `json:"name,omitempty"`
	CreatedDate string            `json:"created_date,omitempty"`
	UpdatedDate string            `json:"updated_date,omitempty"`
}

func (c *Category) parseField(model news_category.Model, parseDate ...bool) {
	c.Model = model
	name := make(map[string]string)
	if model.NameEn != "" {
		name[lang.EnUSCode] = model.NameEn
	}
	if model.NameZh != "" {
		name[lang.ZhCNCode] = model.NameZh
	}
	c.Name = name

	if (len(parseDate) > 0) && parseDate[0] {
		c.CreatedDate = model.CreatedAt.String()
		c.UpdatedDate = model.UpdatedAt.String()
	}
}

func (c *Category) FindById(id int64, parseDate ...bool) error {
	model, err := news_category.FindById(id)
	if err != nil {
		return err
	}
	c.parseField(model, parseDate...)

	return nil
}

func (list *CategoryList) FindLimit(query newsValidator.QueryCategory) error {
	models, err := news_category.FindLimit(query)
	if err != nil {
		return err
	}
	for _, model := range models {
		var c Category
		c.parseField(*model)
		*list = append(*list, &c)
	}

	return nil
}
