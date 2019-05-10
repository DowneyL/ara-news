package news

import (
	"ara-news/components/response"
	"ara-news/controllers"
	"ara-news/helper"
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
	case "BatchDelete":
		nc.ValidJSON(&validators.CategoryIds{})
	case "Update":
		nc.ValidJSON(&validators.NewsCategory{})
	case "UpdateNameEn":
		nc.ValidJSON(&validators.UpdateNameEn{})
	}
}

func (nc *CategoryController) List() {
	categories, err := news_category.FindAll()
	if err != nil {
		nc.ErrorJSON(response.QUERY_ERROR, err.Error())
	}
	if categories == nil {
		nc.SuccessJSON(new(struct{}))
	}

	nc.SuccessJSON(categories)
}

func (nc *CategoryController) Detail() {
	idStr := nc.Ctx.Input.Param(":id")
	id := helper.StringToInt64(idStr)
	category, err := news_category.FindById(id)
	if err != nil {
		nc.ErrorJSON(response.QUERY_ERROR, err.Error())
	}

	nc.SuccessJSON(category)
}

func (nc *CategoryController) Create() {
	var data validators.NewsCategory
	_ = json.Unmarshal(nc.Ctx.Input.RequestBody, &data)
	num, err := news_category.Insert(data)
	if err != nil {
		nc.SystemErrorJSON()
	}

	nc.SuccessJSON(helper.NewAffectNum(num))
}

func (nc *CategoryController) Update() {
	var data validators.NewsCategory
	idStr := nc.Ctx.Input.Param(":id")
	id := helper.StringToInt64(idStr)
	_ = json.Unmarshal(nc.Ctx.Input.RequestBody, &data)
	num, e := news_category.UpdateById(id, data)
	if e != nil {
		nc.ErrorJSON(response.QUERY_ERROR, e.Error())
	}

	nc.SuccessJSON(helper.NewAffectNum(num))
}

func (nc *CategoryController) UpdateNameEn() {
	var data validators.UpdateNameEn
	idStr := nc.Ctx.Input.Param(":id")
	id := helper.StringToInt64(idStr)
	_ = json.Unmarshal(nc.Ctx.Input.RequestBody, &data)
	num, e := news_category.UpdateNameEnById(id, data)
	if e != nil {
		nc.ErrorJSON(response.QUERY_ERROR, e.Error())
	}

	nc.SuccessJSON(helper.NewAffectNum(num))
}

func (nc *CategoryController) Delete() {
	idStr := nc.Ctx.Input.Param(":id")
	id := helper.StringToInt64(idStr)
	num, err := news_category.DeleteById(id)
	if err != nil {
		nc.SystemErrorJSON()
	}

	nc.SuccessJSON(helper.NewAffectNum(num))
}

func (nc *CategoryController) BatchDelete() {
	var cIds validators.CategoryIds
	_ = json.Unmarshal(nc.Ctx.Input.RequestBody, &cIds)
	num, err := news_category.DeleteByIds(cIds)
	if err != nil {
		nc.SystemErrorJSON()
	}

	nc.SuccessJSON(helper.NewAffectNum(num))
}
