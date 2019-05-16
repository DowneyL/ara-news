package news

import (
	"ara-news/controllers"
	"ara-news/helper"
	"ara-news/models/news_category"
	newsService "ara-news/services/news"
	newsValidator "ara-news/validators/news"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type CategoryController struct {
	controllers.BaseController
}

func (nc *CategoryController) BeforeAction() {
	_, action := nc.GetControllerAndAction()
	switch action {
	case "List":
		nc.ValidForm(&newsValidator.QueryCategory{})
	case "Create":
		nc.ValidJSON(&newsValidator.Category{})
	case "BatchDelete":
		nc.ValidJSON(&newsValidator.CategoryIds{})
	case "Update":
		nc.ValidJSON(&newsValidator.Category{})
	case "UpdateNameEn":
		nc.ValidJSON(&newsValidator.UpdateNameEn{})
	}
}

func (nc *CategoryController) List() {
	var (
		query      newsValidator.QueryCategory
		categories newsService.CategoryList
	)
	_ = nc.ParseForm(&query)
	err := categories.FindLimitCategory(query)
	if err != nil {
		nc.QueryErrorJSON(err.Error())
	}
	if categories == nil {
		nc.SuccessJSON(new(struct{}))
	}

	nc.SuccessJSON(categories)
}

func (nc *CategoryController) Detail() {
	var category newsService.Category
	idStr := nc.Ctx.Input.Param(":id")
	id := helper.StringToInt64(idStr)
	err := category.FindCategoryById(id, true)
	if err != nil {
		if err == orm.ErrNoRows {
			nc.SuccessJSON(new(struct{}))
		}
		nc.QueryErrorJSON(err.Error())
	}

	nc.SuccessJSON(category)
}

func (nc *CategoryController) Create() {
	var data newsValidator.Category
	_ = json.Unmarshal(nc.Ctx.Input.RequestBody, &data)
	num, err := news_category.Insert(data)
	if err != nil {
		nc.QueryErrorJSON(err.Error())
	}

	nc.SuccessJSON(helper.NewInsertId(num))
}

func (nc *CategoryController) Update() {
	var data newsValidator.Category
	idStr := nc.Ctx.Input.Param(":id")
	id := helper.StringToInt64(idStr)
	_ = json.Unmarshal(nc.Ctx.Input.RequestBody, &data)
	num, err := news_category.UpdateById(id, data)
	if err != nil {
		nc.QueryErrorJSON(err.Error())
	}

	nc.SuccessJSON(helper.NewAffectNum(num))
}

func (nc *CategoryController) UpdateNameEn() {
	var data newsValidator.UpdateNameEn
	idStr := nc.Ctx.Input.Param(":id")
	id := helper.StringToInt64(idStr)
	_ = json.Unmarshal(nc.Ctx.Input.RequestBody, &data)
	num, err := news_category.UpdateNameEnById(id, data)
	if err != nil {
		nc.QueryErrorJSON(err.Error())
	}

	nc.SuccessJSON(helper.NewAffectNum(num))
}

func (nc *CategoryController) Delete() {
	idStr := nc.Ctx.Input.Param(":id")
	id := helper.StringToInt64(idStr)
	num, err := news_category.DeleteById(id)
	if err != nil {
		nc.SystemErrorJSON(err.Error())
	}

	nc.SuccessJSON(helper.NewAffectNum(num))
}

func (nc *CategoryController) BatchDelete() {
	var cIds newsValidator.CategoryIds
	_ = json.Unmarshal(nc.Ctx.Input.RequestBody, &cIds)
	num, err := news_category.DeleteByIds(cIds)
	if err != nil {
		nc.SystemErrorJSON(err.Error())
	}

	nc.SuccessJSON(helper.NewAffectNum(num))
}
