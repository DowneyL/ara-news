package help

import (
	helpValidator "ara-news/components/validators/help"
	"ara-news/controllers"
	"ara-news/helper"
	"ara-news/models/help_document_category"
	"ara-news/models/help_document_content"
	helpService "ara-news/services/help"
	"encoding/json"
)

type Controller struct {
	controllers.BaseController
}

func (c *Controller) BeforeAction() {
	_, action := c.GetControllerAndAction()
	switch action {
	case "Create":
		c.ValidJSON(&helpValidator.Content{})
	case "List":
		c.ValidForm(&helpValidator.Query{})
	}
}

func (c *Controller) Create() {
	var content helpValidator.Content
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &content)
	exist := help_document_category.Exist(content.Cid)
	if !exist {
		c.InvalidArgumentJSON()
	}
	i, err := help_document_content.Insert(content)
	if err != nil {
		c.QueryErrorJSON(err.Error())
	}

	c.SuccessJSON(helper.NewInsertId(i))
}

func (c *Controller) List() {
	var (
		categories helpService.Categories
		query      helpValidator.Query
	)
	_ = c.ParseForm(&query)
	err := categories.FindLimit(query)
	if err != nil {
		c.QueryErrorJSON(err.Error())
	}

	c.SuccessJSON(categories)
}
