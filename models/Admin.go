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
	CreataTime  int64
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

func UpdateAdmin(admin *Admin) error {
	o := orm.NewOrm()
	err := o.Read(admin)
	if err != nil {
		return err
	}
	_,err = o.Update(admin)
	return err
}
