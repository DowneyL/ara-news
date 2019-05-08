package news

import (
	"ara-news/controllers"
	"ara-news/models/news_category"
	"ara-news/validators"
	"encoding/json"
)

type CategoryController struct {
	controllers.BaseController
}

func (nc *CategoryController) BeforeAction() {
	_, action := nc.GetControllerAndAction()
	switch action {
	case "Create":
		nc.ValidJSON(&validators.NewsCategory{})
	}
}

func (nc *CategoryController) List() {
	categories := news_category.FindAll()
	if categories == nil {
		nc.SystemErrorJSON()
	}

	nc.SuccessJSON(categories)
}

func (nc *CategoryController) Detail() {
	id := nc.Ctx.Input.Param(":id")
	nc.Ctx.WriteString("Category Detail " + id)
}

func (nc *CategoryController) Create() {
	var data validators.NewsCategory
	_ = json.Unmarshal(nc.Ctx.Input.RequestBody, &data)
	if _, err := news_category.Insert(data); err != nil {
		nc.SystemErrorJSON()
	}

	nc.SuccessJSON(new(struct{}))
}
