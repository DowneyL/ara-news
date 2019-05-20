package news

import (
	"ara-news/helper"
	"ara-news/models/news_info"
	newsValidator "ara-news/validators/news"
)

type Info struct {
	news_info.Model
	Platform      string `json:"platform"`
	PublishedDate string `json:"published_date"`
	CreatedDate   string `json:"created_date"`
	UpdatedDate   string `json:"updated_date"`
}

type InfoList []*Info

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

func (il *InfoList) FindLimit(query newsValidator.Query) error {
	models, err := news_info.FindLimit(query)
	if err != nil {
		return err
	}

	for _, model := range models {
		var info Info
		info.parseField(*model)
		*il = append(*il, &info)
	}

	return nil
}

func (il *InfoList) GetAllId() (ids []int64, cIds []int64) {
	for _, info := range *il {
		ids = append(ids, info.Id)
		cIds = append(cIds, info.Cid)
	}
	cIds = helper.RmDuplicate(cIds)

	return
}
