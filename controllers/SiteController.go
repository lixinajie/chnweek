package controllers
import(
	_"github.com/astaxie/beego"
	"chnweek/models"
	"chnweek/utils"
	_"fmt"
	"strings"
	"time"
)

type SiteController struct {
	BaseController
}

func (this *SiteController) Prepare() {
	
}

//登录页面
func (this *SiteController) Login() {
    //判断是否已经登录
	adminInfo := this.GetSession("adminInfo")
	if adminInfo != nil {
		this.Ctx.Redirect(302,"/")
	}
	this.TplName = "site/login.html"	
}

//登录
func (this *SiteController) DoLogin() {
	userName := strings.TrimSpace(this.GetString("userName"))
	password := strings.TrimSpace(this.GetString("password"))
	if userName == "" || password == "" {
			this.jsonResult(1000,"账号、密码不能为空")		
	}
	//根据用户名查询用户信息
	adminInfo,err := models.AdminGetOneByUserName(userName)	
	if err != nil {
		this.jsonResult(1000,"账号、密码有误")
	}	
	if adminInfo.Status != 1 {
		this.jsonResult(1000,"账号不存在")
	}
	if adminInfo.Password != utils.Md5(password) {
		this.jsonResult(1000,"账号、密码有误")
	}
	this.SetSession("adminInfo",adminInfo)
	//登录成功后修改用户登录ip等信息
	adminInfo.LastTime = time.Now().Unix() 
	ip := this.Ctx.Input.IP()	
	adminInfo.LastTime,_ = utils.Ip2Long(ip)
	models.UpdateAdmin(adminInfo)
	this.jsonResult(200,"登录成功")	
}

