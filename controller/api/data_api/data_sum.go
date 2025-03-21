package data_api

import (
	"backend/controller/res"
	"backend/global"
	"backend/service/dataService"
	"github.com/gin-gonic/gin"
)

func (DataApi) DataSumView(c *gin.Context) {
	resp, err := dataService.DataSumService()
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(resp, c)
}
