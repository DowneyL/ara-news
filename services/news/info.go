package news

import "ara-news/models/news_info"

type Info struct {
	news_info.Model
	Platform      string `json:"platform"`
	PublishedDate string `json:"published_date"`
	CreatedDate   string `json:"created_date"`
	UpdatedDate   string `json:"updated_date"`
}

func (i *Info) parseField(info news_info.Model) {
	i.Model = info
	i.Platform = info.AttributeSetId.String()
	i.CreatedDate = info.CreatedAt.String()
	i.UpdatedDate = info.UpdatedAt.String()
	i.PublishedDate = info.PublishedAt.String()
}

func (i *Info) FindById(id int64) error {
	info, err := news_info.FindById(id)
	if err != nil {
		return err
	}
	i.parseField(info)

	return nil
}
