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

func (MessageApi) MessageCreateView(c *gin.Context) {
	var cr req.MessageRequest
	title := "发送消息"
	log := logStash.NewAction(c)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	msg, err := msgService.MsgCreateService(cr, claims.UserID)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}

	log.InfoItem(title, "发送成功", fmt.Sprintf("用户%s给用户%d发送%s", claims.Username, cr.RevUserID, cr.Content))
	res.OkWithMessage(msg, c)
}
