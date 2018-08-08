package sysinit 
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

func initMysqlDatabase(){
	//连接名称
  	dbAlias := beego.AppConfig.String("db_alias")
	//数据库名称
	dbName 	:= beego.AppConfig.String("db_name")
	//用户名
	dbUser 	:= beego.AppConfig.String("db_user")
	//密码
	dbPwd 	:= beego.AppConfig.String("db_pwd")
	//数据库域名（ip）
	dbHost 	:= beego.AppConfig.String("db_host")
	//数据库端口
	dbPort	:= beego.AppConfig.String("db_port")
	//编码
	dbCharset	:= beego.AppConfig.String("db_charset")
	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset, 30)
	//如果是开发模式，则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	//自动建表
	orm.RunSyncdb("default", false, isDev)
	if isDev {
		orm.Debug = isDev
	}
}









