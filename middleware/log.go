package middleware

import (
	"backend/plugins/logStash"
	"bytes"
	"github.com/gin-gonic/gin"
)

type responseWrite struct {
	gin.ResponseWriter
	byteData *bytes.Buffer
}

func (rw responseWrite) Write(buf []byte) (int, error) {
	rw.byteData.Write(buf)
	return rw.ResponseWriter.Write(buf)
}

func LogMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求
		r := responseWrite{
			ResponseWriter: c.Writer,
			byteData:       bytes.NewBuffer([]byte{}),
		}
		c.Writer = r

		// 判断是否为图片上传请求
		contentType := c.GetHeader("Content-Type")
		isImageUpload := contentType == "multipart/form-data"

		if !isImageUpload {
			c.Next()
			// 响应
			_action, ok := c.Get("action")
			if !ok {
				return
			}
			action, ok := _action.(*logStash.Action)
			if !ok {
				return
			}
			action.SetResponseContent(r.byteData.String())
			action.SetFlush()
		} else {
			// 如果是图片上传请求，直接调用后续中间件和处理函数，不记录日志
			c.Next()
		}
	}
}

func InitActionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		action := logStash.NewAction(c)
		action.SetRequestHeader(c)
		c.Set("action", action)
		c.Next()
	}
}
