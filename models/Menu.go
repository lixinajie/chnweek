package models
import(
	"github.com/astaxie/beego/orm"
)
type Menu struct {
	Id 	int64
	Pid	int64
	Level	int64
	Name 	string
	Controller	string
	Method		string
	Icon		string
	IsShow	int8
	Sort 	int64
	Status	int8
	IsPermissions	int8
	CreateTime	int64
	UpdateTime	int64
	Children	[]*Menu `orm:"-"`
}

func(this *Menu) TableName() string {
	return MenuTableName()
}

func MenuTableName() string {
	return TableName("menu")
}

func MenuGetAll() ([]*Menu,error) {
	a := make([]*Menu,0)
	_,err := orm.NewOrm().QueryTable(MenuTableName()).All(&a)
	if err != nil{
		return nil,err
	}
	return a,err
}
