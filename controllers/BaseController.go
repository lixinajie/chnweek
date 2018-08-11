package controllers
import(
	"github.com/astaxie/beego"
	"chnweek/models"
	"strings"
	"fmt"
)

type BaseController struct {
	beego.Controller
	adminInfo *models.Admin
	controllerName string
	actionName string
}

func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	//判断是否已经登录
	//this.checkLogin()
	this.getMenu()
}

//json返回
func (this *BaseController) jsonResult (code int,msg string,data ...interface{}) {
	reponse := make(map[string]interface{})
	reponse["code"] = code
	reponse["msg"] = msg
	reponse["data"] = data
	this.Data["json"] = reponse
	this.ServeJSON()
	this.StopRun()
}

//判断是否已经登录
func (this *BaseController) checkLogin() {
	//adminInfo := this.GetSession("adminInfo")
	adminInfo := []int{1,2}
	if adminInfo == nil{
		this.Ctx.Redirect(302,"/login")
	}
}


//设置模板
func (this *BaseController) setTpl(tpl ...string) {
	var tplName string
	if len(tpl) > 0 {
		tplName = strings.Join([]string{tpl[0], "html"}, ".")
	}	else {
		tplName = this.controllerName + "/" + this.actionName + ".html"
	}
	if this.Data["siteName"] == nil {
		this.Data["siteName"] =	beego.AppConfig.String("web_name") 
	}
	this.Layout = "public/layout.html"
	this.TplName = tplName	
}

//获取菜单
func (this *BaseController) getMenu() {
	menuList,err := models.MenuGetAll()
	if err == nil && len(menuList)>0 {
		fmt.Println(menuList)		
	}
}


