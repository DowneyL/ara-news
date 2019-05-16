package news

import (
	"ara-news/controllers"
	"ara-news/helper"
	"ara-news/models/news_content"
	"ara-news/models/news_info"
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
	case "CreateContent":
		c.ValidJSON(&newsValidator.Content{})
	}
}

func (c *Controller) Create() {
	var data newsValidator.News
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	id, err := newsService.Create(data)
	if err != nil {
		c.QueryErrorJSON(err.Error())
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
		c.QueryErrorJSON(err.Error())
	}

	c.SuccessJSON(newsDetail)
}

func (c *Controller) List() {
	var query newsValidator.Query
	_ = c.ParseForm(&query)
	details, err := newsService.FindLimit(query)
	if err != nil {
		c.QueryErrorJSON(err.Error())
	}

	c.SuccessJSON(details)
}

func (c *Controller) CreateContent() {
	var content newsValidator.Content
	nid := helper.StringToInt64(c.Ctx.Input.Param(":id"))
	exist := news_info.Exist(nid)
	if !exist {
		c.InvalidArgumentJSON()
	}
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &content)
	contentId, err := news_content.Insert(nid, content)
	if err != nil {
		c.QueryErrorJSON(err.Error())
	}

	c.SuccessJSON(helper.NewInsertId(contentId))
}
