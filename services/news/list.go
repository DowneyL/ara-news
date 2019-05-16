package news

import (
	"ara-news/models/news_info"
	newsValidator "ara-news/validators/news"
	"errors"
)

type List []*Detail

func (list *List) GetNIds() []int64 {
	var ids []int64
	for _, detail := range *list {
		ids = append(ids, detail.Id)
	}

	return ids
}

func (list *List) SetLimitContent() error {
	var query newsValidator.Query
	nIds := list.GetNIds()
	if len(nIds) < 1 {
		return errors.New("empty news ids")
	}
	query.Ids = list.GetNIds()
	//models, err := news_content.FindLimit(query)
	//if err != nil {
	//	return err
	//}
	//modelMap := make(map[int64]*news_content.Model, len(models))
	//for _, model := range models {
	//	modelMap[model.Nid] = model
	//}
	//for _, detail := range *list {
	//	detail.Contents = *modelMap[detail.Id]
	//	detail.parseContentField(*modelMap[detail.Id])
	//}

	return nil
}

func (list *List) FindInfoLimit(query newsValidator.Query) error {
	models, err := news_info.FindLimit(query)
	if err != nil {
		return err
	}

	for _, model := range models {
		var d Detail
		d.Model = *model
		d.parseInfoField(*model)
		*list = append(*list, &d)
	}

	return nil
}
