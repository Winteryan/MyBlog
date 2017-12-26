package controllers

import (
	"fmt"
	"hello/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	filters := make([]interface{}, 0)
	filters = append(filters, "type", 1)
	r1, total := models.BlogGetList(1, 4, filters...)
	r2, _ := models.BlogGetList(2, 4, filters...)
	r3, _ := models.BlogGetList(3, 4, filters...)
	fmt.Println(r3[0].Title, total)
	fmt.Println(r3[1].Title, total)
	c.Data["List1"] = r1
	c.Data["List2"] = r2
	c.Data["List3"] = r3
	c.Data["Total"] = total
	c.TplName = "home.html"
}

func (main *MainController) Archive() {
	main.TplName = "archive.html"
}

func (main *MainController) Single() {
	main.TplName = "single.html"
}

func (main *MainController) Contact() {
	main.TplName = "contact.html"
}
