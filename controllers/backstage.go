package controllers

type BackstageController struct {
	BaseController
}

func (c *BackstageController) Index() {
	c.TplName = "backstage/index.html"
}
