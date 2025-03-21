package article_repo

/*package article_repo

import (
	"backend/global"
	"backend/models/esmodels"
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
	"strings"
)

type SearchData struct {
	Key   string `json:"key"`
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
}

// AsyncArticleByFullText 同步文章数据到全文搜索
func AsyncArticleByFullText(id, title, content string) {
	indexList := getSearchIndexDataByContent(id, title, content)

	bulk := global.ESClient.Bulk()
	for _, indexData := range indexList {
		req := elastic.NewBulkIndexRequest().Index(esmodels.FullTextModel{}.Index()).Doc(indexData)
		bulk.Add(req)
	}
	result, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infof("%s 添加成功，共 %d 条", title, len(result.Succeeded()))
}

// DeleteFullTextByArticleID 删除文章数据
func DeleteFullTextByArticleID(id string) {
	boolSearch := elastic.NewTermQuery("key", id)
	res, _ := global.ESClient.
		DeleteByQuery().
		Index(esmodels.FullTextModel{}.Index()).
		Query(boolSearch).
		Do(context.Background())
	logrus.Infof("成功删除 %d 条记录", res.Deleted)
}

func getSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	dataList := strings.Split(content, "\n")
	var isCode bool = false
	var headList, bodyList []string
	var body string
	headList = append(headList, getHeader(title))
	for _, s := range dataList {
		// #{1,6}
		// 判断一下是否是代码块
		if strings.HasPrefix(s, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(s, "#") && !isCode {
			headList = append(headList, getHeader(s))
			bodyList = append(bodyList, getBody(body))
			body = ""
			continue
		}
		body += s
	}
	bodyList = append(bodyList, getBody(body))
	ln := len(headList)
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  id + getSlug(headList[i]),
			Key:   id,
		})
	}
	return searchDataList
}

func getHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return head
}

func getBody(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	return doc.Text()
}

func getSlug(slug string) string {
	return "#" + slug
}

*/
