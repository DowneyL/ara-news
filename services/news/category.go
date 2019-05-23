package news

import (
	"ara-news/boot"
	"ara-news/components/lang"
	"ara-news/models/news_category"
	newsValidator "ara-news/validators/news"
)

type Categories []*Category

type Category struct {
	news_category.Model
	Name        string `json:"name,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
	UpdatedDate string `json:"updated_date,omitempty"`
}

func (c *Category) parseField(model news_category.Model, parseDate ...bool) {
	c.Model = model
	currLang := boot.GetLang()
	switch currLang {
	case lang.EnUSCode:
		c.Name = model.NameEn
	case lang.ZhCNCode:
		c.Name = model.NameZh
	default:
		c.Name = model.NameEn
	}

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

func (list *Categories) FindLimit(query newsValidator.QueryCategory) error {
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

func (list *Categories) ParseToMap() map[int64]*Category {
	categoryMap := make(map[int64]*Category, len(*list))
	for _, category := range *list {
		categoryMap[category.Id] = category
	}

	return categoryMap
}
