package message_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/middleware/jwt"

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
	msg, err := msgService.MsgCreateService(cr, claims.UserID)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
