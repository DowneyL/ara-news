package news

import "ara-news/models/news_content"

func (d *Detail) FindContentByNid(nid int64) error {
	fields := []string{"id", "lang", "title", "content"}
	content, err := news_content.FindByNId(nid, fields...)
	if err != nil {
		return err
	}
	d.Content = content
	d.Content.Lang = content.Lid.String()

	return nil
}
