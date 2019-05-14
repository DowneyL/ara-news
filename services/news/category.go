package news

import "ara-news/models/news_category"

type Category struct {
	news_category.Model
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}

func (d *Detail) FindCategoryByCid(cid int64) error {
	fields := []string{"id", "code", "icon", "name_zh", "name_en"}
	category, err := news_category.FindById(cid, fields...)
	if err != nil {
		return err
	}
	d.Category = category

	return nil
}
