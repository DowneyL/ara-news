package help

import (
	"ara-news/controllers"
	"ara-news/helper"
	"ara-news/models/help_document_category"
	"ara-news/models/help_document_category_content"
	helpService "ara-news/services/help"
	helpValidator "ara-news/validators/help"
	"encoding/json"
)

type CategoryController struct {
	controllers.BaseController
}

func (cc *CategoryController) BeforeAction() {
	_, action := cc.GetControllerAndAction()
	switch action {
	case "Create":
		cc.ValidJSON(&helpValidator.Category{})
	case "CreateContent":
		cc.ValidJSON(&helpValidator.Content{})
	}
}

func (cc *CategoryController) Create() {
	var category helpValidator.Category
	_ = json.Unmarshal(cc.Ctx.Input.RequestBody, &category)
	i, err := helpService.CreateCategory(category)
	if err != nil {
		cc.QueryErrorJSON(err.Error())
	}

	cc.SuccessJSON(helper.NewInsertId(i))
}

func (cc *CategoryController) CreateContent() {
	var content helpValidator.Content
	id := helper.StringToInt64(cc.Ctx.Input.Param(":id"))
	exist := help_document_category.Exist(id)
	if !exist {
		cc.InvalidArgumentJSON()
	}

	_ = json.Unmarshal(cc.Ctx.Input.RequestBody, &content)
	i, err := help_document_category_content.Insert(id, content)
	if err != nil {
		cc.QueryErrorJSON(err.Error())
	}

	cc.SuccessJSON(helper.NewInsertId(i))
}
