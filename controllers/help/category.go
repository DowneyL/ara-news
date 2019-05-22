package help

import (
	"ara-news/controllers"
	"fmt"
)

type CategoryController struct {
	controllers.BaseController
}

func (cc *CategoryController) BeforeAction() {
	_, action := cc.GetControllerAndAction()
	switch action {
	case "Create":
		fmt.Println(action)
	}
}

func (cc *CategoryController) Create() {
	cc.Ctx.WriteString("Create")
}
