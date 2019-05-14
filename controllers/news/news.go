package news

import (
	"ara-news/components/response"
	"ara-news/controllers"
	"ara-news/helper"
	newsService "ara-news/services/news"
	newsValidator "ara-news/validators/news"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type Controller struct {
	controllers.BaseController
}

func (c *Controller) BeforeAction() {
	_, action := c.GetControllerAndAction()
	switch action {
	case "Create":
		c.ValidJSON(&newsValidator.News{})
	case "List":
		c.ValidForm(&newsValidator.Query{})
	}
}

func (c *Controller) Create() {
	var data newsValidator.News
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	id, err := newsService.Create(data)
	if err != nil {
		c.ErrorJSON(response.QUERY_ERROR, err.Error())
	}

	c.SuccessJSON(helper.NewInsertId(id))
}

func (c *Controller) Detail() {
	id := helper.StringToInt64(c.Ctx.Input.Param(":id"))
	newsDetail, err := newsService.FindById(id)
	if err != nil {
		if err == orm.ErrNoRows {
			c.SuccessJSON(new(struct{}))
		}
		c.ErrorJSON(response.QUERY_ERROR, err.Error())
	}

	c.SuccessJSON(newsDetail)
}

func (c *Controller) List() {
	var query newsValidator.Query
	_ = c.ParseForm(&query)
	details, err := newsService.FindLimit(query)
	if err != nil {
		c.ErrorJSON(response.QUERY_ERROR, err.Error())
	}

	c.SuccessJSON(details)
}
