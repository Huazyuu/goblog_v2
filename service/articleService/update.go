package articleService

import (
	"backend/controller/req"
	"backend/global"
	"backend/models/diverseType"
	"backend/models/esmodels"
	"backend/repository/article_repo"
	"backend/repository/img_repo"
	"errors"
	"github.com/fatih/structs"
	"time"
)

func ArticleUpdateService(cr req.ArticleUpdateRequest, userid uint) (string, error) {
	// 鉴权
	idlist, err := article_repo.GetArticleIDListByUserID(userid)
	isexist := false
	for _, id := range idlist {
		if id != cr.ID {
			continue
		} else {
			isexist = true
			break
		}
	}
	if !isexist {
		return "不是本人发布的文章", errors.New("不是本人发布的文章")
	}

	var bannerUrl string
	global.Log.Info(cr.BannerID)
	if cr.BannerID >= 0 {
		bannerUrl, err = img_repo.GetPathByID(cr.BannerID)
		if bannerUrl == "" || err != nil {
			return "banner不存在", errors.New("banner不存在")
		}

	}
	article := esmodels.ArticleModel{
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:     cr.Title,
		Keyword:   cr.Title,
		Abstract:  cr.Abstract,
		Content:   cr.Content,
		Category:  cr.Category,
		Source:    cr.Source,
		Link:      cr.Link,
		BannerID:  cr.BannerID,
		BannerUrl: bannerUrl,
		Tags:      cr.Tags,
	}
	// 封装数据
	maps := structs.Map(&article)
	var DataMap = make(map[string]any)
	// 去掉空值
	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case diverseType.Array:
			if len(val) == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		}
		DataMap[key] = v
	}
	err = article_repo.UpdateArticle(cr.ID, DataMap)
	if err != nil {
		return "更新失败", err
	}
	// todo 同步全文搜索
	return "更新文章成功", nil
}
