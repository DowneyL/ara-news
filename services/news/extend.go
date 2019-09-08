package news

import (
	newsValidator "ara-news/components/validators/news"
	"ara-news/models/news_info_extend"
)

type Extends []*Extend

type Extend struct {
	news_info_extend.Model
}

func (e *Extend) parseFiled(model news_info_extend.Model) {
	e.Model = model
}

func (es *Extends) parseFiled(models []*news_info_extend.Model) {
	for _, model := range models {
		var e Extend
		e.parseFiled(*model)
		*es = append(*es, &e)
	}
}

func (e *Extend) FindByNid(nid int64) error {
	fields := []string{"view_count"}
	extend, err := news_info_extend.FindByNid(nid, fields...)
	if err != nil {
		return err
	}
	e.parseFiled(extend)

	return nil
}

func (es *Extends) FindLimit(query newsValidator.Query) error {
	fields := []string{"nid", "view_count"}
	models, err := news_info_extend.FindLimit(query, fields...)
	if err != nil {
		return err
	}
	es.parseFiled(models)

	return nil
}

func (es *Extends) ParseToMap() map[int64]*Extend {
	extendMap := make(map[int64]*Extend, len(*es))
	for _, extend := range *es {
		extendMap[extend.Nid] = extend
	}

	return extendMap
}
