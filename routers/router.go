package routers

import (
	"chnweek/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login",&controllers.SiteController{},"get:Login")
	beego.Router("/login",&controllers.SiteController{},"post:DoLogin")
}
