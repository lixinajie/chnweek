package models
import (	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Admin),new(Menu))
}

func TableName(name string) string {
	return beego.AppConfig.String("db_prefix")+name
}

