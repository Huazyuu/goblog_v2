package usersService

import (
	"backend/controller/req"
	"backend/controller/resp"
	"backend/middleware/jwt"
	"backend/models/diverseType"
	"backend/models/sqlmodels"
	"strings"
)

func UsersList(claims *jwt.CustomClaims, page req.UserListRequest) ([]resp.UserResponse, int64, error) {
	list, count, err := req.ComList(sqlmodels.UserModel{Role: diverseType.Role(page.Role)}, req.Option{
		PageInfo: page.PageInfo,
		Likes:    []string{"nick_name", "user_name"},
	})
	if err != nil {
		return nil, 0, err
	}
	var users []resp.UserResponse
	for _, user := range list {
		if diverseType.Role(claims.Role) != diverseType.PermissionAdmin {
			// 管理员
			user.UserName = ""
		}
		user.Tel = desensitizationTel(user.Tel)
		user.Email = desensitizationEmail(user.Email)
		// 脱敏
		users = append(users, resp.UserResponse{
			UserModel: user,
			RoleID:    int(user.Role),
		})
	}
	return users, count, nil
}

func desensitizationEmail(email string) string {
	// 256655@qq.com  2****@qq.com
	// yaheb7479@yaho.com  y****@yaho.com
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}

func desensitizationTel(tel string) string {
	// 15852526585
	// 158****6585
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}
