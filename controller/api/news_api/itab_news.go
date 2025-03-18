package news_api

import (
	"backend/controller/res"
	"backend/global"
	"backend/service/redisService"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type params struct {
	ID   string `form:"id"`
	Size int    `form:"size"`
}

type header struct {
	Signaturekey string `header:"signaturekey"`
	Version      string `header:"version"`
	UserAgent    string `header:"user-agent"`
}

type NewResponse struct {
	Code int                    `json:"code"`
	Data []redisService.NewData `json:"data"`
	Msg  string                 `json:"msg"`
}

const (
	newAPI   = "https://api.codelife.cc/api/top/list"
	timeout  = 2 * time.Second
	langCN   = "cn"
	cacheKey = "%s-%d"
)

func (NewsApi) NewListView(c *gin.Context) {
	var cr params
	var headers header

	if err := c.ShouldBindQuery(&cr); err != nil || c.ShouldBindHeader(&headers) != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	if cr.Size < 1 {
		cr.Size = 1
	}

	key := fmt.Sprintf(cacheKey, cr.ID, cr.Size)
	if newsData, _ := redisService.GetNews(key); len(newsData) > 0 {
		res.OkWithData(newsData, c)
		return
	}

	query := url.Values{
		"lang": []string{langCN},
		"id":   []string{cr.ID},
		"size": []string{strconv.Itoa(cr.Size)},
	}

	resp, err := sendRequest(newAPI, query, structs.Map(headers))
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("服务请求失败", c)
		return
	}
	defer resp.Body.Close()

	var response NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("数据解析失败", c)
		return
	}

	if response.Code != http.StatusOK {
		res.FailWithMessage(response.Msg, c)
		return
	}

	res.OkWithData(response.Data, c)
	redisService.SetNews(key, response.Data)
}

func sendRequest(api string, query url.Values, headers map[string]interface{}) (*http.Response, error) {
	client := &http.Client{Timeout: timeout}
	req, _ := http.NewRequest("GET", api+"?"+query.Encode(), nil)

	// 设置固定请求头
	req.Host = "api.codelife.cc"
	for k, v := range map[string]string{
		"Accept":     "*/*",
		"Connection": "keep-alive",
	} {
		req.Header.Set(k, v)
	}

	// 设置动态请求头
	for k, v := range headers {
		req.Header.Set(k, fmt.Sprint(v))
	}

	return client.Do(req)
}
