package chat_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/plugins/logStash"
	"backend/repository/chat_repo"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ChatApi) ChatRemoveView(c *gin.Context) {
	var cr req.RemoveRequest
	title := "删除聊天记录"
	log := logStash.NewAction(c)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	err, cnt := chat_repo.RemoveChatMsg(cr.IDList)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage("群聊记录删除失败", c)
		global.Log.Error(err.Error())
		return
	}
	log.InfoItem(title, "创建文章成功", fmt.Sprintf("共删除记录%d条,list:%T", cnt, cr.IDList))
	res.OkWithMessage(fmt.Sprintf("共删除记录%d条", cnt), c)
}
