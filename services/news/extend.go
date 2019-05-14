package news

import "ara-news/models/news_info_extend"

func (d *Detail) FindExtendByNid(nid int64) error {
	fields := []string{"view_count"}
	extend, err := news_info_extend.FindByNid(nid, fields...)
	if err != nil {
		return err
	}
	d.Extend = extend

	return nil
}
