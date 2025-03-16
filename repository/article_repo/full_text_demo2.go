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
		req := elastic.NewBulkIndexRequest().
			Index(esmodels.FullTextModel{}.Index()).
			Doc(indexData)
		bulk.Add(req)
	}

	// 执行批量操作并处理错误
	result, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}

	// 检查失败项
	if len(result.Failed()) > 0 {
		logrus.Errorf("批量操作中%d项失败", len(result.Failed()))
		for _, fail := range result.Failed() {
			logrus.Errorf("失败项: %+v", fail)
		}
	} else {
		logrus.Infof("%s 添加成功，共 %d 条", title, len(result.Succeeded()))
	}
}

// DeleteFullTextByArticleID 删除文章数据
func DeleteFullTextByArticleID(id string) {
	boolSearch := elastic.NewTermQuery("key", id)
	res, err := global.ESClient.DeleteByQuery().
		Index(esmodels.FullTextModel{}.Index()).
		Query(boolSearch).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infof("成功删除 %d 条记录", res.Deleted)
}

func getSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	dataList := strings.Split(content, "\n")
	var isCode bool
	var headList, bodyList []string
	var body strings.Builder

	// 处理初始标题
	initialTitle := getHeader(title)
	if initialTitle != "" {
		headList = append(headList, initialTitle)
	}

	body.Reset() // 初始化body为空

	for _, line := range dataList {
		// 处理代码块标记
		if strings.HasPrefix(line, "```") {
			isCode = !isCode
			continue // 跳过代码块标记行本身
		}

		// 跳过代码块内容
		if isCode {
			body.WriteString(line + "\n") // 保留代码块内容到正文
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

	// 处理剩余正文内容
	currentBody := body.String()
	bodyList = append(bodyList, getBody(currentBody))

	// 校验标题和正文列表长度
	if len(headList) != len(bodyList) {
		logrus.Warnf("标题列表长度%d，正文列表长度%d，尝试自动修复", len(headList), len(bodyList))
		minLen := min(len(headList), len(bodyList))
		headList = headList[:minLen]
		bodyList = bodyList[:minLen]
	}

	// 生成最终数据
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

// 辅助函数：取较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 优化后的标题处理函数
func getHeader(head string) string {
	head = strings.TrimSpace(strings.TrimPrefix(head, "#"))
	head = strings.TrimSpace(head)
	return head
}

// 优化后的正文处理函数
func getBody(body string) string {
	// 使用扩展生成 HTML 内容
	html := blackfriday.Run(
		[]byte(body),
		blackfriday.WithExtensions(blackfriday.CommonExtensions|
			blackfriday.Tables|
			blackfriday.FencedCode|
			blackfriday.AutoHeadingIDs),
	)

	// 解析 HTML 并提取文本
	doc, _ := goquery.NewDocumentFromReader(bytes.NewBuffer(html))
	text := doc.Text()

	return strings.TrimSpace(text)
}
func getSlug(slug string) string {
	// 生成slug，替换空格为连字符，小写
	slug = strings.ToLower(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-") // 可选，处理下划线
	return "#" + slug
}
