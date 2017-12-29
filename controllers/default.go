package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	filters := make([]interface{}, 0)
	filters = append(filters, "type", "1")
	r1, total := models.BlogGetList(1, 4, filters...)
	r2, _ := models.BlogGetList(2, 4, filters...)
	r3, _ := models.BlogGetList(3, 4, filters...)
	c.Data["List1"] = r1
	c.Data["List2"] = r2
	c.Data["List3"] = r3
	c.Data["Total"] = total
	c.TplName = "portals/home.html"
}

func (main *MainController) Archive() {
	main.TplName = "portals/archive.html"
}

func (main *MainController) Single() {
	main.TplName = "portals/single.html"
}

func (main *MainController) Contact() {
	main.TplName = "portals/contact.html"
}
