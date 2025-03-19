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
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	log := logStash.NewAction(c)

	msg, err := msgService.MsgCreateService(cr, claims.UserID)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}

	log.SetItemInfo("发送消息", fmt.Sprintf("用户%s给用户%d发送消息", claims.Username, cr.RevUserID))
	log.Info(cr.Content)
	res.OkWithMessage(msg, c)
}
