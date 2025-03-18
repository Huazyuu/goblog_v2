package logStash

import (
	"backend/global"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"reflect"
	"strings"
)

type Action struct {
	IP          string    `json:"ip"`       // 客户端IP地址（用于定位来源）
	Addr        string    `json:"addr"`     // 客户端地理位置（通过IP解析）
	UserName    string    `json:"username"` // 登录用户名称（从Token解析）
	ServiceName string    `json:"service"`  // 当前服务名称（如API服务名称）
	UserID      uint      `json:"user_id"`  // 用户ID（从Token解析）
	Level       Level     `json:"level"`    // 日志级别（Info/Warning/Error）
	Title       string    `json:"title"`    // 日志标题（用户自定义的简要描述）
	ItemList    []string  // 日志详细项列表（包含HTML格式的结构化内容）
	Model       *LogModel // 数据库模型对象（用于关联数据库记录）
	Token       string    `json:"token"`    // 用户认证令牌（用于解析用户信息）
	LogType     LogType   `json:"log_type"` // 日志类型（如操作日志、登录日志等）
}

func NewAction(c *gin.Context) *Action {
	ip := c.ClientIP()
	addr := getAddr(ip)
	action := &Action{
		IP:       ip,
		Addr:     addr,
		LogType:  ActionType,
		ItemList: make([]string, 0), // 初始化空切片
	}
	auth := c.GetHeader("Authorization")
	token := splitToken(auth)
	action.SetToken(token)
	return action
}

// Info 记录信息级日志
func (a *Action) Info(title string) {
	a.Level = Info
	a.Title = title
	a.save()
}

// Warn 记录警告级日志
func (a *Action) Warn(title string) {
	a.Level = Warning
	a.Title = title
	a.save()
}

// Error 记录错误级日志
func (a *Action) Error(title string) {
	a.Level = Error
	a.Title = title
	a.save()
}

// SetItemInfo 设置信息级日志项
func (a *Action) SetItemInfo(label string, value interface{}) {
	a.setItem(label, value, Info)
}

// SetItemWarn 设置警告级日志项
func (a *Action) SetItemWarn(label string, value interface{}) {
	a.setItem(label, value, Warning)
}

// SetItemErr 设置错误级日志项
func (a *Action) SetItemErr(label string, value interface{}) {
	a.setItem(label, value, Error)
}

// SetItem 设置通用日志项（默认信息级）
func (a *Action) SetItem(label string, value interface{}) {
	a.setItem(label, value, Info)
}

// SetToken 设置认证令牌
func (a *Action) SetToken(token string) {
	a.Token = token
}

// SetImage 添加图片日志项
func (a *Action) SetImage(url string) {
	a.ItemList = append(a.ItemList, fmt.Sprintf(`<div class="log_image"><img src="%s"></div>`, url))
}

// SetUrl 添加链接日志项
func (a *Action) SetUrl(title, url string) {
	a.ItemList = append(a.ItemList, fmt.Sprintf(`<div class="log_link">
															<a target="_blank" href="%s">%s</a>
														</div>`,
		url, title))
}

// SetUpload 处理文件上传日志
func (a *Action) SetUpload(c *gin.Context) {
	forms, err := c.MultipartForm()
	if err != nil {
		a.SetItem("form错误", err.Error())
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
		a.ItemList = append(a.ItemList, item)
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
	a.ItemList = append(a.ItemList, requestLog)
}

// SetRequestHeader 添加请求头日志项
func (a *Action) SetRequestHeader(c *gin.Context) {
	header := c.Request.Header.Clone()
	jsonData, _ := json.Marshal(header)
	a.ItemList = append(a.ItemList, fmt.Sprintf(`
	<div class="log_request_header">
		<pre>%s</pre>
	</div>`, string(jsonData)))
}

// SetResponse 设置一组出参
func (a *Action) SetResponse(c *gin.Context) {
	c.Set("action", a)
}

// SetResponseContent 记录响应内容
func (a *Action) SetResponseContent(response string) {
	a.ItemList = append(a.ItemList, fmt.Sprintf(`<div class="log_response">%s</div>`, response))
}

// SetFlush 更新日志记录（追加内容）
func (a *Action) SetFlush() {
	if a.Model != nil {
		a.Level = a.Model.Level
	}
	a.save()
}

// save 持久化日志到数据库
func (a *Action) save() {
	content := strings.Join(a.ItemList, "\n")
	if a.Token != "" {
		payload := parseToken(a.Token)
		if payload != nil {
			a.UserID = payload.UserID
			a.UserName = payload.UserName
		}
	}
	if a.Model == nil {
		// 创建新日志记录
		newLog := &LogModel{
			IP:          a.IP,
			Addr:        a.Addr,
			Level:       a.Level,
			Title:       a.Title,
			Content:     content,
			UserID:      a.UserID,
			UserName:    a.UserName,
			ServiceName: a.ServiceName,
			Type:        a.LogType,
		}
		global.DB.Create(newLog)
		a.Model = newLog
		a.ItemList = nil // 清空已提交内容
	} else {
		// 更新现有记录
		global.DB.Model(a.Model).Updates(map[string]interface{}{
			"level":   a.Level,
			"title":   a.Title,
			"content": fmt.Sprintf("%s\n%s", a.Model.Content, content),
		})
	}
}

// setItem 核心日志项设置方法
func (a *Action) setItem(label string, value interface{}, level Level) {
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
		a.ItemList = append(a.ItemList, item)
	default:
		// 基础类型直接格式化
		item := fmt.Sprintf(`
		<div class="log_item %s">
			<div class="log_item_label">%s</div>
			<div class="log_item_content">%v</div>
		</div>`, level.String(), label, value)
		a.ItemList = append(a.ItemList, item)
	}
}
