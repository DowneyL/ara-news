package news

import (
	"ara-news/components/mysql"
	"ara-news/models/news_category"
	"ara-news/models/news_content"
	"ara-news/models/news_info"
	"ara-news/models/news_info_extend"
	newsValidator "ara-news/validators/news"
)

type Detail struct {
	news_info.Model
	Platform      string                 `json:"platform"`
	PublishedDate string                 `json:"published_date"`
	CreatedDate   string                 `json:"created_date"`
	UpdatedDate   string                 `json:"updated_date"`
	Category      news_category.Model    `json:"category"`
	Content       news_content.Model     `json:"content"`
	Extend        news_info_extend.Model `json:"extend"`
}

func CreateNews(news newsValidator.News) (int64, error) {
	o := mysql.GetOrmer("master")
	err := o.Begin()
	info := news.Info
	infoModel, err := news_info.NewModel(info)
	if err != nil {
		return 0, err
	}
	nid, err := o.Insert(&infoModel)
	if err != nil {
		_ = o.Rollback()
		return 0, err
	}

	//TODO:起协程插入
	content := news.Content
	contentModel := news_content.NewModel(nid, content)
	_, err = o.Insert(&contentModel)
	if err != nil {
		_ = o.Rollback()
		return 0, err
	}

	extendModel := news_info_extend.NewModel(nid)
	_, err = o.Insert(&extendModel)
	if err != nil {
		_ = o.Rollback()
		return 0, err
	}
	_ = o.Commit()

	return nid, nil
}

func FindById(id int64) (Detail, error) {
	var detail Detail
	// TODO:起协程查询
	err := detail.FindInfoById(id)
	if err != nil {
		return detail, err
	}
	_ = detail.FindContentByNid(id)
	_ = detail.FindExtendByNid(id)
	_ = detail.FindCategoryByCid(detail.Cid)

	return detail, nil
}
