package news

import (
	"ara-news/components/mysql"
	"ara-news/models/news_content"
	"ara-news/models/news_info"
	"ara-news/models/news_info_extend"
	newsValidator "ara-news/validators/news"
)

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
