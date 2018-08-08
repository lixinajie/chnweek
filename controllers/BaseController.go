package controllers
import(
	"github.com/astaxie/beego"
	"fmt"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	fmt.Println("hello Base")

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













