package news

import (
	"ara-news/models/news_info"
	newsValidator "ara-news/validators/news"
)

func (d *Detail) parseInfoField(info news_info.Model) {
	d.Platform = info.AttributeSetId.String()
	d.CreatedDate = info.CreatedAt.String()
	d.UpdatedDate = info.UpdatedAt.String()
	d.PublishedDate = info.PublishedAt.String()
}

func (d *Detail) FindInfoById(id int64) error {
	info, err := news_info.FindById(id)
	if err != nil {
		return err
	}
	d.Model = info
	d.parseInfoField(info)

	return nil
}

func (list *List) FindInfoLimit(query newsValidator.Query) error {
	models, err := news_info.FindLimit(query)
	if err != nil {
		return err
	}

	for _, model := range models {
		var d Detail
		d.Model = *model
		d.parseInfoField(*model)
		*list = append(*list, &d)
	}

	return nil
}
