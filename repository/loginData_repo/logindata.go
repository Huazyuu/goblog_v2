package loginData_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
)

/*============ database functions  ============*/

func CreateLoginData(l sqlmodels.LoginDataModel) error {
	return global.DB.Create(&l).Error
}
