package news

import "ara-news/models/news_info_extend"

type Extend struct {
	news_info_extend.Model
}

func (e *Extend) FindByNid(nid int64) error {
	fields := []string{"view_count"}
	extend, err := news_info_extend.FindByNid(nid, fields...)
	if err != nil {
		return err
	}
	e.Model = extend

	return nil
}
