package news

import (
	"ara-news/components/mysql"
	"ara-news/models/news_category"
	"ara-news/models/news_content"
	"ara-news/models/news_info"
	"ara-news/models/news_info_extend"
	newsValidator "ara-news/validators/news"
)

type News struct {
	news_info.Model
	Platform string                 `json:"platform"`
	Category news_category.Model    `json:"category"`
	Content  news_content.Model     `json:"content"`
	Extend   news_info_extend.Model `json:"extend"`
}

var news News

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

func FindById(id int64) (News, error) {
	// TODO:起协程查询
	info, err := news_info.FindById(id)
	if err != nil {
		return news, err
	}
	news.Model = info
	news.Platform = info.AttributeSetId.String()

	fields := []string{"id", "lang", "title", "content"}
	content, err := news_content.FindByNId(id, fields...)
	if err != nil {
		return news, err
	}
	news.Content = content

	fields = []string{"view_count"}
	extend, err := news_info_extend.FindByNid(id, fields...)
	if err != nil {
		return news, err
	}
	news.Extend = extend

	fields = []string{"id", "code", "icon", "name_zh", "name_en"}
	category, err := news_category.FindById(info.Cid, fields...)
	if err != nil {
		return news, err
	}
	news.Category = category

	return news, nil
}
