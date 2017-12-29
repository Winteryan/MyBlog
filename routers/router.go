package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/archive", &controllers.MainController{}, "get:Archive")
	beego.Router("/single", &controllers.MainController{}, "get:Single")
	beego.Router("/contact", &controllers.MainController{}, "get:Contact")
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/addArticle", &controllers.AddarticleController{}, "*:Add")
	beego.Router("/listArticle", &controllers.AddarticleController{}, "*:List")
	beego.Router("/updateArticle", &controllers.AddarticleController{}, "*:Update")
	beego.Router("/deleteArticle", &controllers.AddarticleController{}, "*:Delete")
	beego.Router("/lookArticle", &controllers.AddarticleController{}, "*:Look")
	beego.Router("/backstage", &controllers.BackstageController{}, "*:Index")
}
