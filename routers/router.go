package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/archive", &controllers.MainController{}, "*:Archive")
	beego.Router("/single", &controllers.MainController{}, "*:Single")
	beego.Router("/contact", &controllers.MainController{}, "*:Contact")
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/addArticle", &controllers.AddarticleController{}, "*:Add")
	beego.Router("/listArticle", &controllers.AddarticleController{}, "*:List")
	beego.Router("/updateArticle", &controllers.AddarticleController{}, "*:Update")
	beego.Router("/deleteArticle", &controllers.AddarticleController{}, "*:Delete")
	beego.Router("/lookArticle", &controllers.AddarticleController{}, "*:Look")
	beego.Router("/backstage", &controllers.BackstageController{}, "*:Index")
	beego.Router("/banner", &controllers.BannerController{}, "*:Index")
	beego.Router("/updateBanner", &controllers.BannerController{}, "*:Update")
	//	beego.Router("/allblog", &controllers.PortalController{}, "*:Allblog")
}
