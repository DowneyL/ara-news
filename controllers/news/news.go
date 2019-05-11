package news

import (
	"ara-news/components/response"
	"ara-news/controllers"
	"ara-news/helper"
	newsService "ara-news/services/news"
	newsValidator "ara-news/validators/news"
	"encoding/json"
)

type Controller struct {
	controllers.BaseController
}

func (c *Controller) BeforeAction() {
	_, action := c.GetControllerAndAction()
	switch action {
	case "Create":
		c.ValidJSON(&newsValidator.News{})
	}
}

func (c *Controller) Create() {
	var data newsValidator.News
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	id, err := newsService.CreateNews(data)
	if err != nil {
		c.ErrorJSON(response.QUERY_ERROR, err.Error())
	}

	c.SuccessJSON(helper.NewInsertId(id))
}
