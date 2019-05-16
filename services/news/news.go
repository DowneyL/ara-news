package news

import (
	"ara-news/components/mysql"
	"ara-news/models/news_content"
	"ara-news/models/news_info"
	"ara-news/models/news_info_extend"
	newsValidator "ara-news/validators/news"
)

type Detail struct {
	Info
	Category Category `json:"category"`
	Contents Contents `json:"contents"`
	Extend   Extend   `json:"extend"`
}

func Create(news newsValidator.News) (int64, error) {
	o := mysql.GetOrmer("master")
	err := o.Begin()
	nid, err := news_info.TransactionInsert(o, news.Info)
	if err != nil {
		return 0, err
	}
	err = news_content.TransactionInsert(o, nid, news.Content)
	if err != nil {
		return 0, err
	}
	err = news_info_extend.TransactionInsert(o, nid)
	if err != nil {
		return 0, err
	}
	_ = o.Commit()

	return nid, nil
}

func FindById(id int64) (Detail, error) {
	var detail Detail
	// TODO:起协程查询
	err := detail.Info.FindById(id)
	if err != nil {
		return detail, err
	}
	_ = detail.Contents.FindAllByNid(id)
	_ = detail.Extend.FindByNid(id)
	_ = detail.Category.FindById(detail.Info.Cid)

	return detail, nil
}
