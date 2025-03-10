package menu_api

import (
	"backend/models/req"
	"backend/models/res"
	"backend/service/menuService"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (MenuApi) MenuDetailByIDView(c *gin.Context) {
	id := c.Param("id")
	uint64id, _ := strconv.ParseUint(id, 10, 64)
	respData, err := menuService.GetMenuDetail(uint(uint64id))
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(respData, c)
}
func (MenuApi) MenuDetailByPathView(c *gin.Context) {
	var cr req.MenuDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	respData, err := menuService.GetMenuDetailByPath(cr.Path)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(respData, c)

}
