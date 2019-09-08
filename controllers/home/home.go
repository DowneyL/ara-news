package home

import "ara-news/controllers"

type Controller struct {
	controllers.BaseController
}

func (c *Controller) Index() {
	var data map[string]interface{}
	var banner []string

	imageHost := "http://192.168.1.107:8080/static/assets/images/"
	banner = append(banner,
		imageHost+"banner/volkswagen.jpg",
		imageHost+"banner/aston_martin.jpg",
		imageHost+"banner/delorean.jpg",
		imageHost+"banner/seckill.jpg",
	)

	data = make(map[string]interface{})
	data["banner"] = banner

	c.SuccessJSON(data)
}
