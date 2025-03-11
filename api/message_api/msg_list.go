package message_api

import (
	"backend/global"
	"backend/models/common"
	"backend/models/req"
	"backend/models/res"
	"backend/models/sqlmodels"
	"backend/service/msgService"
	"github.com/gin-gonic/gin"
)

// MessageListAllView 所有消息列表
func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr common.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, cnt, err := msgService.MsgAllList(cr)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithList(list, cnt, c)
}

// MessageUserListView 某人收到不同发送者消息的列表
func (MessageApi) MessageUserListView(c *gin.Context) {
	var cr req.MessageUserListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithMessage("参数绑定错误", c)
		return
	}
	list, cnt, err := msgService.MessageUserList(cr)
	if err != nil {
		res.FailWithMessage("获取数据失败", c)
		return
	}
	res.OkWithList(list, cnt, c)

}

// MessageUserListByUserView 某个用户的聊天列表，包含与该用户有聊天记录的其他用户信息以及他们之间的消息数量
func (MessageApi) MessageUserListByUserView(c *gin.Context) {
	var cr req.MessageUserListByUserRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMessage("参数错误", c)
		return
	}
	list, err := msgService.MessageUserListByUser(cr.UserID)
	if err != nil || len(list) == 0 {
		res.FailWithMessage("查询失败,没有对应用户", c)
		return
	}
	res.OkWithList(list, int64(len(list)), c)

}

// MessageUserRecordView 两个用户之间的聊天列表
func (MessageApi) MessageUserRecordView(c *gin.Context) {
	var cr req.MessageUserRecordRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMessage("请传入对应的id", c)
		return
	}
	list, count, _ := common.ComList(sqlmodels.MessageModel{}, common.Option{
		PageInfo: cr.PageInfo,
		Where:    global.DB.Where("(send_user_id = ? and rev_user_id = ? ) or ( rev_user_id = ? and send_user_id = ? )", cr.SendUserID, cr.RevUserID, cr.SendUserID, cr.RevUserID),
	})
	res.OkWithList(list, count, c)
}
