package msgService

import (
	"backend/global"
	"backend/models/req"
	"backend/repository/msg_repo"
	"errors"
	"fmt"
)

func MessageRemove(cr req.RemoveRequest, uid uint, role int) (string, error) {
	// 管理员
	if role == 1 {
		mlist, err := msg_repo.GetMsgsByIDList(cr.IDList)
		if err != nil {
			global.Log.Error(err.Error())
			return "查找出错", err
		}
		if len(mlist) == 0 {
			return "消息不存在", errors.New("消息不存在")
		}
		cnt, err := msg_repo.DeleteMsgs(mlist)
		if err != nil {
			global.Log.Error(err.Error())
			return "删除失败", err
		}
		return "删除成功 " + fmt.Sprintf("共删除 %d 个消息", cnt), err

	} else if role == 2 {
		mlist, err := msg_repo.GetMyMsgsByIDList(uid, cr.IDList)
		if err != nil {
			global.Log.Error(err.Error())
			return "查找出错", err
		}
		if len(mlist) == 0 {
			return "消息不存在或非本人发送的消息", errors.New("消息不存在")
		}

		cnt, err := msg_repo.DeleteMsgs(mlist)
		if err != nil {
			global.Log.Error(err.Error())
			return "删除失败", err
		}
		return "删除成功 " + fmt.Sprintf("共删除 %d 个消息", cnt), err
	}
	return "没有权限", nil
}
