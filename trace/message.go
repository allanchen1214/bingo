package trace

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	KeyMessage   = "TraceMsg"
	KeyStartTime = "StartTime" // 后端首次接收到请求时赋值，后端之间的服务调用需把该值拷贝到上下文，而非重新赋值
	KeySequence  = "Sequence"
	KeyUserID    = "UserID"
	KeyUsername  = "Username"
	KeyClientIP  = "ClientIP"
	KeyDeviceID  = "DeviceID"
)

// Message 跟踪上下文结构定义
type Message struct {
	Context   context.Context
	StartTime time.Time
	Logger    *zap.Logger

	Sequence string
	UserID   string
	Username string
	ClientIP string
	DeviceID string

	ExtData map[string]interface{}
}

// GinMessage 初始化设置或获取Gin Http Server跟踪上下文
func GinMessage(c *gin.Context) *Message {
	val, ok := c.Get(KeyMessage)
	var msg *Message
	if ok {
		msg, ok = val.(*Message)
		if ok {
			return msg
		}
	}
	msg = &Message{Context: c, StartTime: time.Now()}
	c.Set(KeyMessage, msg)
	return msg
}
