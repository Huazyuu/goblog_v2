package chat_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/repository/chat_repo"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ChatApi) ChatRemoveView(c *gin.Context) {
	var cr req.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	err, cnt := chat_repo.RemoveChatMsg(cr.IDList)
	if err != nil {
		res.FailWithMessage("群聊记录删除失败", c)
		global.Log.Error(err.Error())
		return
	}
	res.OkWithMessage(fmt.Sprintf("共删除记录%d条", cnt), c)
}
