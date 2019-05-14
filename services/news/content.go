package news

import "ara-news/models/news_content"

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
