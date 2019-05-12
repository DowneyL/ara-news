package news

import "ara-news/models/news_category"

type Category struct {
	news_category.Model
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}
