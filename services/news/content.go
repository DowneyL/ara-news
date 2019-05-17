package news

import (
	"ara-news/models/news_content"
	newsValidator "ara-news/validators/news"
)

type Contents []*Content

type Content struct {
	news_content.Model
	Lang string `json:"lang,omitempty"`
}

func (c *Content) parseField(model news_content.Model) {
	c.Model = model
	c.Lang = model.LangType.String()
}

func (cs *Contents) parseField(models []*news_content.Model) {
	for _, model := range models {
		var c Content
		c.parseField(*model)
		*cs = append(*cs, &c)
	}
}

func (cs *Contents) FindAllByNid(nid int64) error {
	fields := []string{"id", "lang_type", "title", "content", "is_default"}
	models, err := news_content.FindAllByNId(nid, fields...)
	if err != nil {
		return err
	}
	cs.parseField(models)

	return nil
}

func (cs *Contents) FindLimit(query newsValidator.Query) error {
	fields := []string{"id", "nid", "lang_type", "title", "content", "is_default"}
	models, err := news_content.FindLimit(query, fields...)
	if err != nil {
		return err
	}
	cs.parseField(models)

	return nil
}

func (cs *Contents) ParseToMap() map[int64]*Content {
	contentMap := make(map[int64]*Content, len(*cs))
	for _, content := range *cs {
		contentMap[content.Nid] = content
	}

	return contentMap
}
