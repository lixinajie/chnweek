package controllers
import(
	"github.com/astaxie/beego"
	_"fmt"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	//判断是否已经登录
}

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
	adminInfo := this.GetSession("adminInfo")
	if adminInfo == nil{
		this.Ctx.Redirect(302,"/login")
	}
}











