package data_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/service/dataService"
	"github.com/gin-gonic/gin"
)

func (DataApi) LoginDataView(c *gin.Context) {
	var cr req.DateRequest
	_ = c.ShouldBindQuery(&cr)
	resp, err := dataService.LoginDataService(cr)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithData(resp, c)
}
