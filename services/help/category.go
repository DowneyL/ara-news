package help

import (
	helpValidator "ara-news/components/validators/help"
	"ara-news/models/help_document_category"
)

type Categories []*Category

type Category struct {
	help_document_category.Model
	Platform    string `json:"platform"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}

func (c *Category) parseField(model help_document_category.Model) {
	c.Model = model
	c.Platform = model.AttributeSetId.String()
	c.CreatedDate = model.CreatedAt.String()
	c.UpdatedDate = model.UpdatedAt.String()
}

func (cs *Categories) FindLimit(query helpValidator.Query) error {
	models, err := help_document_category.FindLimit(query)
	if err != nil {
		return err
	}
	for _, model := range models {
		var c Category
		c.parseField(*model)
		*cs = append(*cs, &c)
	}

	return nil
}
