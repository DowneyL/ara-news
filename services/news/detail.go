package news

import (
	"ara-news/models/news_info"
	"ara-news/models/news_info_extend"
)

type Detail struct {
	news_info.Model
	Platform      string                 `json:"platform"`
	PublishedDate string                 `json:"published_date"`
	CreatedDate   string                 `json:"created_date"`
	UpdatedDate   string                 `json:"updated_date"`
	Category      Category               `json:"category"`
	Contents      Contents               `json:"contents"`
	Extend        news_info_extend.Model `json:"extend"`
}

func (d *Detail) parseInfoField(info news_info.Model) {
	d.Model = info
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
	d.parseInfoField(info)

	return nil
}

func (d *Detail) FindExtendByNid(nid int64) error {
	fields := []string{"view_count"}
	extend, err := news_info_extend.FindByNid(nid, fields...)
	if err != nil {
		return err
	}
	d.Extend = extend

	return nil
}

func (d *Detail) FindContentByNid(nid int64) error {
	var contents Contents
	err := contents.FindAllByNid(nid)
	d.Contents = make(Contents, 0)
	if err != nil {
		return err
	}
	if contents != nil {
		d.Contents = contents
	}

	return nil
}

func (d *Detail) FindCategoryByCid(cid int64) error {
	var category Category
	err := category.FindById(cid)
	if err != nil {
		return err
	}
	d.Category = category

	return nil
}
