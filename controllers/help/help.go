package help

import (
	"ara-news/controllers"
	"ara-news/helper"
	"ara-news/models/help_document_category"
	"ara-news/models/help_document_content"
	helpValidator "ara-news/validators/help"
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
