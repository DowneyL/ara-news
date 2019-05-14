package news

import "ara-news/models/news_info"

func (d *Detail) FindInfoById(id int64) error {
	info, err := news_info.FindById(id)
	if err != nil {
		return err
	}
	d.Model = info
	d.Platform = info.AttributeSetId.String()
	d.CreatedDate = info.CreatedAt.String()
	d.UpdatedDate = info.UpdatedAt.String()
	d.PublishedDate = info.PublishedAt.String()

	return nil
}
