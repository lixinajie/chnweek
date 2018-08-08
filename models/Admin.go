package models
import(
	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id int64
	UserName string
	RealName string
	Password string
	Mail     string
	Phone    string
	LastTime int64
	LastIp 	 int64
	Status   int8
	CretaTime  int64
	UpdateTime int64
}

func (this *Admin) TableName()  string{
	return AdminTableName()
}

func AdminTableName() string {
	return TableName("admin")
}

func AdminGetOneByUserName(userName string) (*Admin,error) {
	a := new(Admin)
	err := orm.NewOrm().QueryTable(AdminTableName()).Filter("user_name",userName).One(a)
	if err != nil{
		return nil,err
	}
	return a,nil
}

