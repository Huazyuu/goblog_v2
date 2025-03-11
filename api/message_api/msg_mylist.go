package message_api

import (
	"backend/global"
	"backend/middleware/jwt"
	"backend/models/common"
	"backend/models/req"
	"backend/models/res"
	"backend/models/sqlmodels"
	"backend/service/msgService"
	"fmt"
	"github.com/gin-gonic/gin"
)

// MessageListView 与自己相关的消息列表
func (MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	list, err := msgService.MsgList(claims)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(list, c)
}

// MessageUserListByMeView 我与其他用户的聊天列表
func (m MessageApi) MessageUserListByMeView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	c.Request.URL.RawQuery = fmt.Sprintf("userID=%d", claims.UserID)
	m.MessageUserListByUserView(c)
}

// MessageUserRecordByMeView 我与某个用户的聊天列表
func (m MessageApi) MessageUserRecordByMeView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr req.MessageUserRecordByMeRequest
	c.ShouldBindQuery(&cr)

	cr.Sort = "created_at asc"
	list, count, _ := common.ComList(sqlmodels.MessageModel{}, common.Option{
		PageInfo: cr.PageInfo,
		Where:    global.DB.Where("(send_user_id = ? and rev_user_id = ? ) or ( rev_user_id = ? and send_user_id = ? )", claims.UserID, cr.UserID, claims.UserID, cr.UserID),
	})

	res.OkWithList(list, count, c)
}
