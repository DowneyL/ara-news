package news

import (
	"ara-news/controllers"
	"ara-news/models/news_category"
	"ara-news/validators"
	"encoding/json"
	"strconv"
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
	}
}

func (nc *CategoryController) List() {
	categories := news_category.FindAll()
	if categories == nil {
		nc.SuccessJSON(new(struct{}))
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

func (nc *CategoryController) Delete() {
	idStr := nc.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	num, err := news_category.DeleteById(id)
	if err != nil {
		nc.SystemErrorJSON()
	}
	result := make(map[string]int64)
	result["count"] = num

	nc.SuccessJSON(result)
}

func (nc *CategoryController) BatchDelete() {
	var cids validators.CategoryIds
	_ = json.Unmarshal(nc.Ctx.Input.RequestBody, &cids)
	num, err := news_category.DeleteByIds(cids)
	if err != nil {
		nc.SystemErrorJSON()
	}
	result := make(map[string]int64)
	result["count"] = num

	nc.SuccessJSON(result)
}
