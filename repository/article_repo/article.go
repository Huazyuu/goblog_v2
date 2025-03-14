package article_repo

import (
	"backend/global"
	"backend/models/esmodels"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"strings"
)

func CreateArticle(article *esmodels.ArticleModel) error {
	indexResponse, err := global.ESClient.Index().
		Index(article.Index()).
		BodyJson(article).Refresh("true").Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	article.ID = indexResponse.Id
	return nil
}

func ISExistData(article esmodels.ArticleModel) bool {
	res, err := global.ESClient.
		Search(article.Index()).
		Query(elastic.NewTermQuery("keyword", article.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

func GetArticleByID(id string) (article esmodels.ArticleModel, err error) {
	res, err := global.ESClient.
		Get().
		Index(article.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return article, err
	}
	err = json.Unmarshal(res.Source, &article)
	if err != nil {
		return
	}
	article.ID = res.Id
	// todo look digg comment count
	// article.LookCount =article.LookCount + redis_ser.NewArticleLook().Get(res.Id)
	// article.DiggCount =article.DiggCount + redis_ser.NewDigg().Get(res.Id)
	// article.CommentCount =article.CommentCount + redis_ser.NewCommentCount().Get(res.Id)
	return article, err
}

func GetArticleIDListByUserID(userid uint) (articleIDList []string, err error) {
	// 构建 Elasticsearch 查询条件，使用 TermQuery 来精确匹配 userid 字段
	query := elastic.NewTermQuery("user_id", userid)
	result, err := global.ESClient.Search().
		Index(esmodels.ArticleModel{}.Index()). // 指定索引名称
		Query(query).                           // 设置查询条件
		Size(10000).                            // 设置返回结果的最大数量，可根据实际情况调整
		Do(context.Background())                // 执行搜索请求
	if err != nil {
		return nil, fmt.Errorf("failed to search articles by user ID: %w", err)
	}

	// 遍历搜索结果，提取文章 ID
	for _, hit := range result.Hits.Hits {
		articleIDList = append(articleIDList, hit.Id)
	}

	return articleIDList, nil
}

func GetArticleIDList() (articleIDList []string, err error) {
	result, err := global.ESClient.Search().
		Index(esmodels.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to search articles by user ID: %w", err)
	}
	for _, hit := range result.Hits.Hits {
		articleIDList = append(articleIDList, hit.Id)
	}
	return articleIDList, nil
}

func GetArticleByKeyword(keyword string) (article esmodels.ArticleModel, err error) {
	res, err := global.ESClient.Search().
		Index(article.Index()).
		Query(elastic.NewTermQuery("keyword", keyword)).
		Size(1).
		Do(context.Background())
	if err != nil {
		return
	}
	if res.Hits.TotalHits.Value == 0 {
		return article, errors.New("文章不存在")
	}
	hit := res.Hits.Hits[0]

	err = json.Unmarshal(hit.Source, &article)
	if err != nil {
		return
	}
	article.ID = hit.Id
	return
}

func UpdateArticle(id string, data map[string]any) error {
	_, err := global.ESClient.
		Update().
		Index(esmodels.ArticleModel{}.Index()).
		Id(id).
		Doc(data).Refresh("true").
		Do(context.Background())
	return err
}

func GetArticleList(option Option) (list []esmodels.ArticleModel, count int, err error) {
	if option.Query == nil {
		option.Query = elastic.NewBoolQuery()
	}
	if option.Key != "" {
		option.Query.Must(elastic.NewMultiMatchQuery(option.Key, option.Fields...))
	}
	if option.Category != "" {
		option.Query.Must(elastic.NewMultiMatchQuery(option.Category, "category"))
	}
	if option.Tag != "" {
		option.Query.Must(elastic.NewMultiMatchQuery(option.Tag, "tags"))
	}

	sortField := struct {
		field string
		asc   bool
	}{
		field: "created_at",
		asc:   false, // 从小到大  从大到小
	}
	if option.Sort != "" {
		splist := strings.Split(option.Sort, " ")
		field := splist[0]
		way := splist[1]
		if len(splist) == 2 {
			if way == "asc" || way == "desc" {
				sortField.field = field
				if way == "asc" {
					sortField.asc = true
				}
				if way == "desc" {
					sortField.asc = false
				}
			}
		}
	}

	res, err := global.ESClient.
		Search(esmodels.ArticleModel{}.Index()).
		Query(option.Query).
		Highlight(elastic.NewHighlight().Field("title")).
		From(option.GetForm()).
		Sort(sortField.field, sortField.asc).
		Size(option.Limit).
		Do(context.Background())
	if err != nil {
		return
	}
	// type res.SearchHits.[]*SearchHit
	/*
		前面的 Hits 是 SearchHits 结构体的字段名，
		后面的 Hits 是 SearchHits 结构体里的一个切片字段，类型为 []*SearchHit。
		这个切片中存储了每一个具体匹配到的文档的信息，
		每个元素是一个指向 SearchHit 结构体的指针，
		SearchHit 结构体包含了单个匹配文档的详细信息，
		例如文档的 ID、文档的源数据等。
	*/
	count = int(res.Hits.TotalHits.Value)

	// 返回值处理
	// todo diggInfo lookInfo commentInfo
	for _, hit := range res.Hits.Hits {
		var resp esmodels.ArticleModel
		err = json.Unmarshal(hit.Source, &resp)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		title, ok := hit.Highlight["title"]
		if ok {
			resp.Title = title[0]
		}
		resp.ID = hit.Id
		// todo diggInfo lookInfo commentInfo count
		// resp.DiggCount = resp.DiggCount + digg
		// resp.LookCount = resp.LookCount + look
		// resp.CommentCount = resp.CommentCount + comment
		list = append(list, resp)
	}
	return list, count, nil
}

func RemoveArticleByList(idlist []string) (int, error) {
	bulk := global.ESClient.Bulk().Index(esmodels.ArticleModel{}.Index()).Refresh("true")
	for _, id := range idlist {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulk.Add(req)
		// todo 删除全文搜索
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		return 0, err
	}
	return len(res.Succeeded()), nil
}
