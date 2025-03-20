package log_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/service/logService"
	"github.com/gin-gonic/gin"
)

func (LogApi) LogListView(c *gin.Context) {
	var cr req.LogListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, msg, err := logService.LogListService(cr)
	if err != nil {
		res.FailWithMessage(msg, c)
		global.Log.Error(err)
		return
	}
	res.OkWithList(list, int64(len(list)), c)
}
