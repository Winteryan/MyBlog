package controllers

import (
	"fmt"
	"hello/models"
	html "html/template"
	"math"
	"strconv"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	filters := make([]interface{}, 0)
	filters2 := make([]interface{}, 0)
	filters = append(filters, "type", "original")
	r1, total := models.BlogGetList(1, 4, filters...)
	r2, _ := models.BlogGetList(2, 4, filters...)
	r3, _ := models.BlogGetList(3, 4, filters...)
	fmt.Println(r2)
	fmt.Println(r3)
	banners, _ := models.BannerGetList(1, 4, filters2...)
	c.Data["List1"] = r1
	c.Data["List2"] = r2
	c.Data["List3"] = r3
	c.Data["Banners"] = banners
	c.Data["Total"] = total
	c.TplName = "portals/home.html"
}

func (c *MainController) Archive() {
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	filters := make([]interface{}, 0)
	r1, total := models.BlogGetList(page*3-2, 4, filters...)
	r2, _ := models.BlogGetList(page*3-1, 4, filters...)
	r3, _ := models.BlogGetList(page*3, 4, filters...)
	fmt.Println(r2)
	fmt.Println(r3)
	pages := (int)(math.Ceil(float64(total) / float64(12)))
	filters2 := make([]interface{}, 0)
	for a := 1; a <= pages; a++ {
		var tempStr = "<a href=\"/archive?page=" + strconv.Itoa(a) + "\">" + strconv.Itoa(a) + "</a>"
		filters2 = append(filters2, html.HTML(tempStr))
	}
	fmt.Println(filters2)
	c.Data["Pages"] = filters2
	c.Data["List1"] = r1
	c.Data["List2"] = r2
	c.Data["List3"] = r3
	c.Data["Total"] = total

	c.TplName = "portals/archive.html"
}

func (main *MainController) Single() {
	id, _ := main.GetInt("id")
	if id == 0 {
		main.Redirect(beego.URLFor("MainController.Archive"), 302)
		main.StopRun()
	}
	blog, err := models.GetBlogById(id)
	if err != nil {
		main.Redirect(beego.URLFor("MainController.Archive"), 302)
		main.StopRun()
	}
	filters := make([]interface{}, 0)
	filters = append(filters, "catalogid", blog.Catalogid)
	r1, _ := models.BlogGetList(1, 3, filters...)
	main.Data["Related"] = r1
	main.Data["Blog"] = blog
	main.TplName = "portals/single.html"
}

func (main *MainController) Contact() {
	main.TplName = "portals/contact.html"
}
