package news

import "ara-news/models/news_content"

type Contents []*Content

type Content struct {
	news_content.Model
	Lang string `json:"lang"`
}

func (c *Content) parseField(model news_content.Model) {
	c.Model = model
	c.Lang = model.LangType.String()
}

func (cs *Contents) FindAllByNid(nid int64) error {
	fields := []string{"id", "lang_type", "title", "content", "is_default"}
	models, err := news_content.FindAllByNId(nid, fields...)
	if err != nil {
		return err
	}

	for _, model := range models {
		var c Content
		c.parseField(*model)
		*cs = append(*cs, &c)
	}

	return nil
}
