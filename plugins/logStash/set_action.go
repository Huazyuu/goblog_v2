package logStash

import (
	"backend/global"
	"backend/models/diverseType"
	"backend/models/sqlmodels"
	"backend/service/redisService"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"reflect"
	"strings"
)

// Action 日志操作结构体
type Action struct {
	ip          string               // 客户端ip地址
	addr        string               // 客户端地理位置
	userName    string               // 用户名
	serviceName string               // 服务名称
	userID      uint                 // 用户ID
	level       diverseType.LogLevel // 日志级别
	title       string               // 日志标题
	itemList    []string             // 日志条目
	token       string               // 认证令牌
	logType     diverseType.LogType  // 日志类型
	model       *sqlmodels.LogModel  // 创建之后赋值给它，用于后期更新
}

// NewAction 构造函数
func NewAction(c *gin.Context) Action {
	action := Action{
		ip:       c.ClientIP(),
		addr:     getAddr(c.ClientIP()),
		logType:  diverseType.ActionType,
		itemList: make([]string, 0),
		token:    splitToken(c.GetHeader("Authorization")),
	}
	action.setResponse(c)
	return action
}

// SetItemInfo 设置信息级日志项
func (a *Action) SetItemInfo(label string, value interface{}) {
	a.setItem(label, value, diverseType.Info)
}

// SetItemWarn 设置警告级日志项
func (a *Action) SetItemWarn(label string, value interface{}) {
	a.setItem(label, value, diverseType.Warning)
}

// SetItemErr 设置错误级日志项
func (a *Action) SetItemErr(label string, value interface{}) {
	a.setItem(label, value, diverseType.Error)
}

// SetToken 设置认证令牌
func (a *Action) SetToken(token string) {
	a.token = token
}

// SetImage 添加图片日志项
func (a *Action) SetImage(url string) {
	a.itemList = append(a.itemList, fmt.Sprintf(`<div class="log_image"><img src="%s"></div>`, url))
}

// SetUrl 添加链接日志项
func (a *Action) SetUrl(title, url string) {
	a.itemList = append(a.itemList, fmt.Sprintf(`<div class="log_link">
															<a target="_blank" href="%s">%s</a>
														</div>`,
		url, title))
}

// SetUpload 处理文件上传日志
func (a *Action) SetUpload(c *gin.Context) {
	forms, err := c.MultipartForm()
	if err != nil {
		a.SetItemErr("form错误", err.Error())
		return
	}
	for key, headers := range forms.File {
		file := headers[0]
		item := fmt.Sprintf(`
		<div class="log_upload">
			<span>%s</span>
			<span>%s</span>
			<span>%d bytes</span>
		</div>`, key, file.Filename, file.Size)
		a.itemList = append(a.itemList, item)
	}
}

// SetRequest 记录请求信息
func (a *Action) SetRequest(c *gin.Context) {
	method := c.Request.Method
	path := c.Request.URL.String()
	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // 重置请求体

	requestLog := fmt.Sprintf(`
	<div class="log_request">
		<span>%s %s</span>
		<pre>%s</pre>
	</div>`, method, path, string(body))
	a.itemList = append(a.itemList, requestLog)
}

// SetRequestHeader 添加请求头日志项
func (a *Action) SetRequestHeader(c *gin.Context) {
	header := c.Request.Header.Clone()
	jsonData, _ := json.Marshal(header)
	a.itemList = append(a.itemList, fmt.Sprintf(`
	<div class="log_request_header">
		<pre>%s</pre>
	</div>`, string(jsonData)))
}

// SetResponseContent 记录响应内容
func (a *Action) SetResponseContent(response string) {
	a.itemList = append(a.itemList, fmt.Sprintf(`<div class="log_response">%s</div>`, response))
}

// ===================手动保存===================

// Info 记录信息级日志
func (a *Action) Info(title string) {
	a.level = diverseType.Info
	a.title = title
	a.save()
}

// Warn 记录警告级日志
func (a *Action) Warn(title string) {
	a.level = diverseType.Warning
	a.title = title
	a.save()
}

// Error 记录错误级日志
func (a *Action) Error(title string) {
	a.level = diverseType.Error
	a.title = title
	a.save()
}

// SetFlush 更新日志记录（追加内容）
func (a *Action) SetFlush() {
	if a.model != nil {
		a.level = a.model.Level
	}
	a.save()
}

// ===================内部方法===================

// save 日志持久化方法
func (a *Action) save() {
	content := strings.Join(a.itemList, "\n")

	if a.token != "" {
		if payload := parseToken(a.token); payload != nil {
			a.userID = payload.UserID
			a.userName = payload.UserName
		}
	}

	if a.model == nil {
		a.model = &sqlmodels.LogModel{
			IP:          a.ip,
			Addr:        a.addr,
			Level:       a.level,
			Title:       a.title,
			Content:     content,
			UserID:      a.userID,
			UserName:    a.userName,
			ServiceName: a.serviceName,
			Type:        a.logType,
		}
		// 如果不对content进行置空，那么content会重复
		a.itemList = []string{}
		return
	} else {
		a.model.Level = a.level
		a.model.Title = a.title
		a.model.Content = a.model.Content + "\n" + content
	}

	global.Log.Info("------------------", a.model.Type)

	rs := redisService.NewRedisStore(global.Redis, global.DB)
	if err := rs.AddLog(a.model); err != nil {
		global.Log.Error("日志写入Redis失败",
			"error", err.Error(),
			"title", a.title,
			"user", a.userName)
	}
	rs.Flush()
	a.itemList = nil
}

// setItem 核心日志项设置方法
func (a *Action) setItem(label string, value interface{}, level diverseType.LogLevel) {
	// 判断值类型
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice:
		// 复杂类型转JSON并格式化输出
		jsonData, _ := json.MarshalIndent(value, "", "  ")
		item := fmt.Sprintf(`
			<div class="log_item %s">
				<div class="log_item_label">%s</div>
				<div class="log_item_content">%s</div>
			</div>`,
			level.String(), label, string(jsonData))
		a.itemList = append(a.itemList, item)
	default:
		// 基础类型直接格式化
		item := fmt.Sprintf(`
		<div class="log_item %s">
			<div class="log_item_label">%s</div>
			<div class="log_item_content">%v</div>
		</div>`, level.String(), label, value)
		a.itemList = append(a.itemList, item)
	}
}

// setResponse 设置日志
func (a *Action) setResponse(c *gin.Context) {
	c.Set("action", a)
}
