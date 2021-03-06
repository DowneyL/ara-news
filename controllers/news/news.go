package news

import (
	newsValidator "ara-news/components/validators/news"
	"ara-news/controllers"
	"ara-news/helper"
	"ara-news/models/news_content"
	"ara-news/models/news_info"
	newsService "ara-news/services/news"
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
	// 通过此接口创建的内容为默认多语言内容
	data.Content.Default = true
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
	_ = newsService.IncrViewCount(id)

	c.SuccessJSON(newsDetail)
}

func (c *Controller) List() {
	var query newsValidator.Query
	_ = c.ParseForm(&query)
	list, err := newsService.FindLimit(query)
	if err != nil {
		if err == orm.ErrNoRows {
			c.SuccessJSON(struct{}{})
		}
		c.QueryErrorJSON(err.Error())
	}

	c.SuccessJSON(list)
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

func (c *Controller) Delete() {
	id := helper.StringToInt64(c.Ctx.Input.Param(":id"))
	err := newsService.DeleteById(id)
	if err != nil {
		c.QueryErrorJSON(err.Error())
	}

	c.SuccessJSON(struct{}{})
}
