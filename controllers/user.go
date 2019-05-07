package controllers

import (
	"ara-news/components/mysql"
	"ara-news/helper"
	"ara-news/models/user"
	"ara-news/validators"
	"strconv"
)

type UserController struct {
	BaseController
}

func (own *UserController) BeforeAction() {
	_, action := own.GetControllerAndAction()
	switch action {
	case "Create":
		own.ValidForm(&validators.UserRegister{})
	}
}

func (own *UserController) GetAll() {
	userModels, err := user.FindAll()
	if err != nil {
		own.SuccessJSON(new([]interface{}))
	}

	own.SuccessJSON(userModels)
}

func (own *UserController) Get() {
	id, _ := strconv.Atoi(own.Ctx.Input.Param(":id"))
	userModel, err := user.FindOneById(id)
	if err != nil {
		own.SuccessJSON(new(struct{}))
	}

	own.SuccessJSON(userModel)
}

func (own *UserController) Create() {
	var model user.Model
	name := own.GetString("name")
	model.Name = name
	model.Type = 1
	model.CreateAt = helper.Date("Y-m-d H:i:s")
	model.UpdateAt = helper.Date("Y-m-d H:i:s")
	o := mysql.GetOrmer("master")
	id, err := o.Insert(&model)
	if err != nil {
		own.SystemErrorJSON()
	}

	own.SuccessJSON(id)
}
