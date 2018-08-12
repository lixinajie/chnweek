package routers

import (
	"chnweek/controllers"
	"github.com/astaxie/beego"
)

func init() {
        beego.Router("/", &controllers.MainController{})
	beego.Router("/login",&controllers.SiteController{},"get:Login")
	beego.Router("/login",&controllers.SiteController{},"post:DoLogin")
        beego.Router("/index",&controllers.MainController{},"get:Index")
	//menu
        beego.Router("/menu/index",&controllers.MenuController{},"get:Index")
        beego.Router("/menu/test",&controllers.MenuController{},"get:Test")
        beego.Router("/menu/add",&controllers.MenuController{},"get,post:Add")
        beego.Router("/menu/del",&controllers.MenuController{},"post:Del")
}
