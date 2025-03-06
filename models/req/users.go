package req

import (
	"backend/models/common"
	"backend/models/diverseType"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

type UserCreateRequest struct {
	NickName string           `json:"nick_name" binding:"required" msg:"请输入昵称"`  // 昵称
	UserName string           `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	Password string           `json:"password" binding:"required" msg:"请输入密码"`   // 密码
	Role     diverseType.Role `json:"role" binding:"required" msg:"请选择权限"`       // 权限  1 管理员  2 普通用户  3 游客
}

type UserRoleRequest struct {
	Role     diverseType.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NickName string           `json:"nick_name"` // 防止用户昵称非法，管理员有能力修改
	UserID   uint             `json:"user_id" binding:"required" msg:"用户id错误"`
}

type UserUpdatePwdRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"` // 旧密码
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`     // 新密码
}

type UserUpdateInfoRequest struct {
	NickName string `json:"nick_name" structs:"nick_name"`
	Sign     string `json:"sign" structs:"sign"`
	Link     string `json:"link" structs:"link"`
	Avatar   string `json:"avatar" structs:"avatar"`
}
type UserListRequest struct {
	common.PageInfo
	Role int `json:"role" form:"role"`
}
