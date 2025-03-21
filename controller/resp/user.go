package resp

import "backend/models/sqlmodels"

type UserResponse struct {
	sqlmodels.UserModel
	RoleID int `json:"role_id"`
}
