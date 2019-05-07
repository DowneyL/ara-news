package controllers

import (
	"ara/validators"
	"fmt"
)

type HelloController struct {
	BaseController
}

func (own *HelloController) BeforeAction() {
	_, action := own.GetControllerAndAction()
	switch action {
	case "GET":
		own.ValidForm(&validators.User{})
	}
}

func (own *HelloController) Get() {
	name := own.GetString("name")
	helloWord := fmt.Sprintf("Hello %s, I'm ara.", name)
	own.Ctx.WriteString(helloWord)
}
