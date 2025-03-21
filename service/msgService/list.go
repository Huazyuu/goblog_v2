package msgService

import (
	"backend/controller/req"
	"backend/controller/resp"

	"backend/models/sqlmodels"
	"backend/repository/msg_repo"
	"backend/repository/user_repo"
)

func MsgAllList(cr req.PageInfo) ([]sqlmodels.MessageModel, int64, error) {
	list, cnt, err := req.ComList(sqlmodels.MessageModel{}, req.Option{
		PageInfo: cr,
	})
	if err != nil {
		return nil, 0, err
	}
	return list, cnt, nil
}

func MessageUserList(cr req.MessageUserListRequest) ([]resp.MessageUserListResponse, int64, error) {
	// 获取总数
	total, err := msg_repo.CountMessageUsers(cr.NickName)
	if err != nil {
		return nil, 0, err
	}
	// 获取统计信息
	stats, err := msg_repo.GetMessageUserStats(cr.NickName, cr.Page, cr.Limit)
	if err != nil {
		return nil, 0, err
	}
	// 收集用户id
	userIDs := make([]uint, 0, len(stats))
	for _, stat := range stats {
		userIDs = append(userIDs, stat.SendUserID)
	}
	// 获取用户信息
	users, err := user_repo.GetUsersByIDList(userIDs)
	if err != nil {
		return nil, 0, err
	}
	// 构建用户映射
	userMap := make(map[uint]sqlmodels.UserModel)
	for _, user := range users {
		userMap[user.ID] = user
	}
	// 组装响应
	result := make([]resp.MessageUserListResponse, 0, len(stats))
	for _, stat := range stats {
		user, exists := userMap[stat.SendUserID]
		if !exists {
			continue
		}
		result = append(result, resp.MessageUserListResponse{
			UserName: user.UserName,
			NickName: user.NickName,
			UserID:   user.ID,
			Avatar:   user.Avatar,
			Count:    stat.Count,
		})
	}
	return result, total, nil

}

// MessageUserListByUser 获取某个用户的聊天列表，包含与该用户有聊天记录的其他用户信息以及他们之间的消息数量
func MessageUserListByUser(userID uint) ([]resp.MessageUserListResponse, error) {
	// 获取统计信息
	stats, err := msg_repo.GetUserMessageStats(userID)
	if err != nil {
		return nil, err
	}
	// 合并统计结果
	userCounts := make(map[uint]int)
	for _, stat := range stats {
		// 处理发送方 用户是接受方
		if stat.SendUserID != userID {
			userCounts[stat.SendUserID] += stat.Count
		}
		// 处理接收方 用户是发送方
		if stat.RevUserID != userID {
			userCounts[stat.RevUserID] += stat.Count
		}
	}
	// 收集用户ID
	userIDs := make([]uint, 0, len(userCounts))
	for uid := range userCounts {
		userIDs = append(userIDs, uid)
	}
	// 获取用户信息
	users, err := user_repo.GetUsersByIDList(userIDs)
	if err != nil {
		return nil, err
	}
	// 构建用户映射
	userMap := make(map[uint]sqlmodels.UserModel)
	for _, user := range users {
		userMap[user.ID] = user
	}

	result := make([]resp.MessageUserListResponse, 0, len(userCounts))
	for uid, count := range userCounts {
		user, exists := userMap[uid]
		if !exists {
			continue
		}
		result = append(result, resp.MessageUserListResponse{
			UserName: user.UserName,
			NickName: user.NickName,
			UserID:   user.ID,
			Avatar:   user.Avatar,
			Count:    count,
		})
	}
	return result, nil
}
