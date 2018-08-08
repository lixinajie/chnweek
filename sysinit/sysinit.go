package sysinit 
import (
	_"chnweek/models"	
	"github.com/astaxie/beego"
)

func init() {
	//开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//初始化数据库
	initMysqlDatabase()
}
