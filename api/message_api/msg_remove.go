package message_api

import (
	"backend/middleware/jwt"
	"backend/models/req"
	"backend/models/res"
	"backend/service/msgService"
	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageRemoveView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var cr req.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	msg, err := msgService.MessageRemove(cr, claims.UserID, claims.Role)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
