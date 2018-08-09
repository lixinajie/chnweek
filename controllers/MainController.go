package controllers

import (
	_"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func(this *MainController) Prepare() {
	this.BaseController.Prepare()
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
