package services
import(
	"chnweek/models"
)

type MenuService struct {
		
}

func MenuRecursion(menuList []*models.Menu,pid int64) []*models.Menu {
	menuTree := make([]*models.Menu,0)
	if len(menuList)<=0 {
		return menuTree
	}
	for _,v := range menuList {
		if v.Pid == pid {
			v.Children = append(v.Children,MenuRecursion(menuList,v.Id)...)
			menuTree = append(menuTree,v)
		}
	}	
	return menuTree	
}



