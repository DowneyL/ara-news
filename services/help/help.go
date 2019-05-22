package help

import (
	"ara-news/components/mysql"
	"ara-news/models/help_document_category"
	"ara-news/models/help_document_category_content"
	helpValidator "ara-news/validators/help"
)

func CreateCategory(category helpValidator.Category) (int64, error) {
	o := mysql.GetOrmer("master")
	_ = o.Begin()
	id, err := help_document_category.TransactionInsert(o, category)
	if err != nil {
		return 0, err
	}
	_, err = help_document_category_content.TransactionInsert(o, id, category.CateContent)
	if err != nil {
		return 0, err
	}
	err = o.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}
