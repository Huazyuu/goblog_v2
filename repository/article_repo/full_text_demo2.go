package article_repo

import (
	"bytes"
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday/v2"
	"strings"

	"github.com/sirupsen/logrus"

	"backend/global"
	"backend/models/esmodels"
)

// SearchData 定义了用于全文搜索的数据结构
// Key 表示文章的唯一标识
// Body 存储文章的正文内容
// Slug 是包含文章 ID 的跳转地址，用于定位文章中的特定部分
// Title 是文章或文章某部分的标题
type SearchData struct {
	Key   string `json:"key"`
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
}

// AsyncArticleByFullText 函数用于将文章数据同步到全文搜索系统（Elasticsearch）
func AsyncArticleByFullText(id, title, content string) {
	indexList := getSearchIndexDataByContent(id, title, content)
	bulk := global.ESClient.Bulk()
	for _, indexData := range indexList {
		req := elastic.NewBulkIndexRequest().
			Index(esmodels.FullTextModel{}.Index()). // 指定要索引的目标索引
			Doc(indexData)                           // 设置要索引的文档数据
		bulk.Add(req)
	}

	result, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	if len(result.Failed()) > 0 {
		logrus.Errorf("批量操作中%d项失败", len(result.Failed()))
		for _, fail := range result.Failed() {
			logrus.Errorf("失败项: %+v", fail)
		}
	} else {
		logrus.Infof("%s 添加成功，共 %d 条", title, len(result.Succeeded()))
	}
}

// DeleteFullTextByArticleID 函数用于根据文章 ID 从全文搜索系统中删除相应的文章数据
func DeleteFullTextByArticleID(id string) {
	// 创建一个 Elasticsearch 的术语查询对象，用于查找 key 字段等于指定文章 ID 的文档
	boolSearch := elastic.NewTermQuery("key", id)
	res, err := global.ESClient.DeleteByQuery().
		Index(esmodels.FullTextModel{}.Index()). // 指定要操作的索引
		Query(boolSearch).                       // 设置查询条件
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infof("成功删除 %d 条记录", res.Deleted)
}

// getSearchIndexDataByContent 函数用于将文章内容处理成适合全文搜索的索引数据列表
func getSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	// 将文章内容按行分割成字符串切片
	dataList := strings.Split(content, "\n")
	var isCode bool                 // 标记当前是否处于代码块中
	var headList, bodyList []string // 分别存储文章的标题列表和正文列表
	var body strings.Builder        // 用于临时存储当前积累的正文内容

	initialTitle := getHeader(title)
	if initialTitle != "" {
		headList = append(headList, initialTitle)
	}

	body.Reset()

	// 遍历文章内容的每一行
	for _, line := range dataList {
		// 处理代码块标记
		if strings.HasPrefix(line, "```") {
			isCode = !isCode
			continue
		}

		// 如果当前处于代码块中，将代码块内容添加到正文缓冲区
		if isCode {
			body.WriteString(line + "\n")
			continue
		}

		// 检测标题行
		if strings.HasPrefix(line, "#") {
			// 先处理当前积累的正文内容
			currentBody := body.String()
			bodyList = append(bodyList, getBody(currentBody))
			body.Reset()

			// 处理新标题
			header := getHeader(line)
			if header != "" {
				headList = append(headList, header)
			}
			continue
		}

		// 积累正文内容
		body.WriteString(line + "\n")
	}

	// 处理剩余的正文内容
	currentBody := body.String()
	bodyList = append(bodyList, getBody(currentBody))

	// 校验标题和正文列表的长度是否一致
	if len(headList) != len(bodyList) {
		logrus.Warnf("标题列表长度%d，正文列表长度%d，尝试自动修复", len(headList), len(bodyList))
		minLen := min(len(headList), len(bodyList))
		headList = headList[:minLen]
		bodyList = bodyList[:minLen]
	}
	for i := range headList {
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  id + getSlug(headList[i]),
			Key:   id,
		})
	}

	return searchDataList
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// getHeader 函数用于处理标题，去除标题开头的所有 # 字符和前后的空格
func getHeader(head string) string {
	// 从标题左侧移除所有连续的 # 字符
	head = strings.TrimSpace(strings.TrimLeft(head, "#"))
	// 去除标题前后的空格
	head = strings.TrimSpace(head)

	global.Log.Info(head) // 记录处理后的标题信息
	return head
}

// getBody 函数用于处理正文内容，将 Markdown 格式的正文转换为纯文本
func getBody(body string) string {
	// 使用 blackfriday 库将 Markdown 内容转换为 HTML 内容
	html := blackfriday.Run(
		[]byte(body),
		blackfriday.WithExtensions(blackfriday.CommonExtensions|
			blackfriday.Tables|
			blackfriday.FencedCode|
			blackfriday.AutoHeadingIDs),
	)

	doc, _ := goquery.NewDocumentFromReader(bytes.NewBuffer(html))
	text := doc.Text()

	return strings.TrimSpace(text)
}

// getSlug 函数用于生成文章的跳转地址，将标题转换为适合 URL 的格式
func getSlug(slug string) string {
	slug = strings.ToLower(slug)
	// 将标题中的空格替换为连字符
	slug = strings.ReplaceAll(slug, " ", "-")
	// 可选操作，将标题中的下划线替换为连字符
	slug = strings.ReplaceAll(slug, "_", "-")
	return "#" + slug // 添加 # 前缀，形成跳转地址
}
