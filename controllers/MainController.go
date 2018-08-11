package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func(this *MainController) Prepare() {
	this.BaseController.Prepare()
}

func (this *MainController) Get() {
	this.Data["siteName"] = beego.AppConfig.String("web_name")+"管理系统"
	this.TplName = "public/main.html"
}

//系统首页
func (this *MainController)	Index() {
	this.Data["pageTitle"] = "首页"
	this.setTpl("main/index")
}
