package message_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/middleware/jwt"
	"backend/plugins/logStash"
	"fmt"

	"backend/service/msgService"
	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageRemoveView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	title := "删除消息"
	log := logStash.NewAction(c)

	var cr req.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	msg, err := msgService.MessageRemove(cr, claims.UserID, claims.Role)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "删除成功", fmt.Sprintf("用户%s删除文章id:%T", claims.Username, cr.IDList))
	res.OkWithMessage(msg, c)
}
