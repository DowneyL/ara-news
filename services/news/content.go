package news

import (
	"ara-news/models/news_content"
	newsValidator "ara-news/validators/news"
	"errors"
)

func (d *Detail) parseContentField(content news_content.Model) {
	d.Content.LangStr = content.Lid.String()
}

func (d *Detail) FindContentByNid(nid int64) error {
	fields := []string{"id", "lang", "title", "content"}
	content, err := news_content.FindByNId(nid, fields...)
	if err != nil {
		return err
	}
	d.Content = content
	d.parseContentField(content)

	return nil
}

func (list *List) SetLimitContent() error {
	var query newsValidator.Query
	nIds := list.GetNIds()
	if len(nIds) < 1 {
		return errors.New("empty news ids")
	}
	query.Ids = list.GetNIds()
	models, err := news_content.FindLimit(query)
	if err != nil {
		return err
	}
	modelMap := make(map[int64]*news_content.Model, len(models))
	for _, model := range models {
		modelMap[model.Nid] = model
	}
	for _, detail := range *list {
		detail.Content = *modelMap[detail.Id]
		detail.parseContentField(*modelMap[detail.Id])
	}

	return nil
}
