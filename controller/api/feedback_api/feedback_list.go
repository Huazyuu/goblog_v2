package feedback_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/middleware/jwt"
	"backend/models/sqlmodels"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func (FeedbackApi) FeedBackListView(c *gin.Context) {
	var cr req.PageInfo
	c.ShouldBindQuery(&cr)

	var isAdmin bool

	list, count, _ := req.ComList(&sqlmodels.FeedbackModel{}, req.Option{
		PageInfo: cr,
		Likes:    []string{"email", "content"},
	})

	token := c.GetHeader("Authorization")
	claims, err := jwt.ParseToken(utils.SplitToken(token))
	if err == nil {
		if claims.Role == 1 {
			isAdmin = true
		}
	}
	// 如果是普通用户和游客，则显示邮箱的第一位及后缀
	if !isAdmin {
		for _, model := range list {
			model.Email = encryEmail(model.Email)
		}
	}

	res.OkWithList(list, count, c)
}
func encryEmail(email string) string {
	// 256655@qq.com  2****@qq.com
	// yaheb7479@yaho.com  y****@yaho.com
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}
