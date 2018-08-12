package controllers
import(
	"chnweek/models"
	_"chnweek/utils"
	"chnweek/services"
	"fmt"
) 

type MenuController struct {
	BaseController
}

func(this *MenuController) Prepare() {
	this.BaseController.Prepare()
}

func (this *MenuController) Index() {
	menuList,err := models.MenuGetAll()
	if err == nil && len(menuList) > 0 {
		menuList = services.MenuRecursion(menuList,0)
	}	
	this.Data["menuList"] = menuList
	this.setTpl("menu/index")
}

func (this *MenuController) Add() {	
	menuList,err := models.MenuGetAll()
	if err == nil && len(menuList) > 0 {
		menuList = services.MenuRecursion(menuList,0)
	}	
	this.Data["menuList"] = menuList
	this.setTpl("menu/add")
}

func (this *MenuController) Del() {
	id, _ := this.GetInt("id")	
	menuInfo,err := models.MenuGetOneById(int64(id))
	if err != nil {
		this.jsonResult(1000,"不存在")	
	}
	if menuInfo.Status == -1 {
		this.jsonResult(200,"删除成功")
	}
	//查询是否还有子集
		
	fmt.Println(menuInfo)
}

	
func (this *MenuController) Test() {
}
