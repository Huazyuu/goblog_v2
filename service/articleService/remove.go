package articleService

import (
	"backend/controller/req"
	"backend/middleware/jwt"
	"backend/repository/article_repo"
	"errors"
	"fmt"
)

func ArticleRemoveService(cr req.ESIDListRequest, claims *jwt.CustomClaims) (string, error) {
	if claims.Role != 1 {
		// 普通用户只能删除自己的文章
		idlist, err := article_repo.GetArticleIDListByUserID(claims.UserID)
		if err != nil {
			return "系统错误 err:" + err.Error(), err
		}
		isexist := existListAInB(cr.IDList, idlist)
		if !isexist {
			return "非本人发布的文章,没有权限删除", errors.New("非本人发布的文章,没有权限删除")
		}
	}

	cnt, err := article_repo.RemoveArticleByList(cr.IDList)
	if err != nil {
		return "系统错误 err:" + err.Error(), err
	}
	if cnt == 0 {
		return "文章id传入错误", errors.New("传入文章id错误")
	}

	return fmt.Sprintf("删除成功,共删除%d篇文章", cnt), nil

}
func existListAInB(a, b []string) bool {
	bSet := make(map[string]bool, len(b))
	for _, s := range b {
		bSet[s] = true
	}

	for _, s := range a {
		if bSet[s] {
			return true
		}
	}
	return false
}
