package articleService

import (
	"backend/controller/req"
	"backend/global"
	"backend/middleware/jwt"
	"backend/models/esmodels"
	"backend/models/sqlmodels"
	"backend/repository/article_repo"
	"backend/repository/img_repo"
	"backend/repository/tag_repo"
	"backend/repository/user_repo"
	"errors"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"math/rand"
	"strings"
	"time"
)

func ArticleCreateService(cr req.ArticleRequest, claims *jwt.CustomClaims) (string, error) {
	userId := claims.UserID
	userNickName := claims.NickName

	// 处理content xss攻击
	unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	// 是不是有script标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		// 有script标签
		doc.Find("script").Remove()
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		cr.Content = markdown
	}
	// 若没有传入简介就截取正文
	if cr.Abstract == "" {
		abs := []rune(doc.Text())
		// 将content转为html，并且过滤xss，以及获取中文内容
		if len(abs) > 100 {
			cr.Abstract = string(abs[:100])
		} else {
			cr.Abstract = string(abs)
		}
	}
	// 若没有传banner,则随机
	if cr.BannerID == 0 {
		ids, _ := img_repo.GetBannersID()
		if len(ids) == 0 {
			return "没有这张banner图片,请先上传", errors.New("没有banner图片")
		}
		cr.BannerID = ids[rand.Intn(len(ids))]
	}
	bannerUrl, err := img_repo.GetPathByID(cr.BannerID)
	if bannerUrl == "" || err != nil {
		return "banner不存在", errors.New("banner不存在")
	}

	avatar, err := user_repo.GetAvatarByID(userId)
	if err != nil {
		return "用户不存在", errors.New("用户不存在")
	}

	for _, tag := range cr.Tags {
		ok, _ := tag_repo.IsExistTagByTitle(tag)
		if !ok {
			tag_repo.CreateTag(sqlmodels.TagModel{
				MODEL: sqlmodels.MODEL{},
				Title: tag,
			})
		}
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	article := esmodels.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Keyword:      cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       userId,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}

	ok := article_repo.ISExistData(article)
	if ok {
		return "文章已存在", errors.New("文章已存在")
	}
	err = article_repo.CreateArticle(&article)
	if err != nil {
		global.Log.Error(err)
		return "创建失败", errors.New("创建失败")
	}

	go article_repo.AsyncArticleByFullText(article.ID, article.Title, article.Content)

	return fmt.Sprintf("文章[%s]创建成功", cr.Title), nil

}
