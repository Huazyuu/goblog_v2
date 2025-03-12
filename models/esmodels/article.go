package esmodels

import (
	"backend/global"
	"backend/models/diverseType"
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

type ArticleModel struct {
	ID        string `json:"id" structs:"id"`                 // es的id
	CreatedAt string `json:"created_at" structs:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at" structs:"updated_at"` // 更新时间

	Title    string `json:"title" structs:"title"`                // 文章标题
	Keyword  string `json:"keyword,omit(list)" structs:"keyword"` // 关键字
	Abstract string `json:"abstract" structs:"abstract"`          // 文章简介
	Content  string `json:"content,omit(list)" structs:"content"` // 文章内容

	LookCount     int `json:"look_count" structs:"look_count"`         // 浏览量
	CommentCount  int `json:"comment_count" structs:"comment_count"`   // 评论量
	DiggCount     int `json:"digg_count" structs:"digg_count"`         // 点赞量
	CollectsCount int `json:"collects_count" structs:"collects_count"` // 收藏量

	UserID       uint   `json:"user_id" structs:"user_id"`               // 用户id
	UserNickName string `json:"user_nick_name" structs:"user_nick_name"` // 用户昵称
	UserAvatar   string `json:"user_avatar" structs:"user_avatar"`       // 用户头像

	Category string `json:"category" structs:"category"` // 文章分类
	Source   string `json:"source" structs:"source"`     // 文章来源
	Link     string `json:"link" structs:"link"`         // 原文链接

	BannerID  uint   `json:"banner_id" structs:"banner_id"`   // 文章封面id
	BannerUrl string `json:"banner_url" structs:"banner_url"` // 文章封面

	Tags diverseType.Array `json:"tags" structs:"tags"` // 文章标签
}

func (a ArticleModel) Index() string {
	return "article_index"
}

func (a ArticleModel) Mapping() string {
	path := "models/esmodels/article_mapper.json"
	txt, err := os.ReadFile(path)
	if err != nil {
		logrus.Error(err)
		return err.Error()
	}
	return string(txt)
}

func (a ArticleModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(a.Index()).Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return exists
	}
	return exists
}

func (a ArticleModel) CreateIndex() error {
	if a.IndexExists() {
		a.RemoveIndex()
	}
	createIdx, err := global.ESClient.
		CreateIndex(a.Index()).
		BodyString(a.Mapping()).Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !createIdx.Acknowledged {
		logrus.Error("创建失败")
		return err
	}
	logrus.Infof("索引 %s 创建成功", a.Index())
	return nil
}

func (a ArticleModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("删除索引失败")
		return err
	}
	logrus.Info("索引删除成功")
	return nil
}
