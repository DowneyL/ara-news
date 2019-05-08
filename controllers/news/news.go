package news

import "ara-news/controllers"

type Controller struct {
	controllers.BaseController
}

func (nc *Controller) BeforeAction() {
	//_, action := nc.GetControllerAndAction()
	//switch action {
	//case "Create":
	//	nc.ValidJSON()
	//}
}
