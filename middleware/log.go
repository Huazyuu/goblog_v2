package middleware

import (
	"backend/global"
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
		c.Next()
		// 响应
		_action, ok := c.Get("action")
		if !ok {
			global.Log.Info("no action")
			return
		}
		action, ok := _action.(*logStash.Action)
		if !ok {
			return
		}
		action.SetResponseContent(r.byteData.String())
		action.SetFlush()
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
