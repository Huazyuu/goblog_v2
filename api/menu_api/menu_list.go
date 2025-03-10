package menu_api

import (
	"backend/models/res"
	"backend/service/menuService"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuListView(c *gin.Context) {
	menus, err := menuService.GetFullMenuList()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithList(menus, int64(len(menus)), c)
}
func (MenuApi) MenuNameListView(c *gin.Context) {
	menus, err := menuService.GetMenuNameList()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(menus, c)
}
