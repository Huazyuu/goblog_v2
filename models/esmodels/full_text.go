package esmodels

import (
	"backend/global"
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

type FullTextModel struct {
	ID    string `json:"id" structs:"id"`       // es的id
	Key   string `json:"key"`                   // 文章关联的id
	Title string `json:"title" structs:"title"` // 文章标题
	Slug  string `json:"slug" structs:"slug"`   // 标题的跳转地址
	Body  string `json:"body" structs:"body"`   // 文章内容
}

func (f FullTextModel) Index() string {
	return "full_text_index"
}

func (f FullTextModel) Mapping() string {
	path := "models/esmodels/fulltext_mapper.json"
	txt, err := os.ReadFile(path)
	if err != nil {
		logrus.Error(err)
		return err.Error()
	}
	return string(txt)
}

func (f FullTextModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(f.Index()).Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return exists
	}
	return exists
}

func (f FullTextModel) CreateIndex() error {
	if f.IndexExists() {
		f.RemoveIndex()
	}
	createIdx, err := global.ESClient.
		CreateIndex(f.Index()).
		BodyString(f.Mapping()).Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !createIdx.Acknowledged {
		logrus.Error("创建失败")
		return err
	}
	logrus.Infof("索引 %s 创建成功", f.Index())
	return nil
}

func (f FullTextModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(f.Index()).Do(context.Background())
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
