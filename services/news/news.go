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

type ListDetail struct {
	Info
	Category Category `json:"category"`
	Content  Content  `json:"content"`
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

func FindLimit(query newsValidator.Query) ([]*ListDetail, error) {
	var (
		infoList   InfoList
		contents   Contents
		extends    Extends
		categories Categories
		list       []*ListDetail
	)
	err := infoList.FindLimit(query)
	if err != nil {
		return list, err
	}

	ids, cids := infoList.GetAllId()
	query.Ids = ids
	_ = contents.FindLimit(query)
	_ = extends.FindLimit(query)
	queryCategory := newsValidator.QueryCategory{CIds: cids}
	_ = categories.FindLimit(queryCategory)

	contentMap := contents.ParseToMap()
	extendMap := extends.ParseToMap()
	categoryMap := categories.ParseToMap()

	for _, info := range infoList {
		var ld ListDetail
		ld.Info = *info
		if content, ok := contentMap[info.Id]; ok {
			ld.Content = *content
		}
		if extend, ok := extendMap[info.Id]; ok {
			ld.Extend = *extend
		}
		if category, ok := categoryMap[info.Cid]; ok {
			ld.Category = *category
		}

		list = append(list, &ld)
	}

	return list, nil
}

func DeleteById(id int64) error {
	o := mysql.GetOrmer("master")
	err := o.Begin()
	_, err = news_info.TransactionDeleteById(o, id)
	if err != nil {
		return err
	}
	_, err = news_content.TransactionDeleteByNid(o, id)
	if err != nil {
		return err
	}
	_, err = news_info_extend.TransactionDeleteByNid(o, id)
	if err != nil {
		return err
	}
	_ = o.Commit()

	return nil
}
