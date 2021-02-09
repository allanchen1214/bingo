package trace

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	KeyMessage   = "TraceMsg"
	keyStartTime = "StartTime" // 后端首次接收到请求时赋值，后端之间的服务调用需把该值拷贝到上下文，而非重新赋值
	keySequence  = "Sequence"
	keyUserID    = "UserID"
	keyUsername  = "Username"
	keyClientIP  = "ClientIP"
	keyDeviceID  = "DeviceID"
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

func (m *Message) ExtraFields() []zapcore.Field {
	extraFields := []zapcore.Field{
		zap.String(keySequence, m.Sequence),
		zap.Time(keyStartTime, m.StartTime),
		zap.String(keyUserID, m.UserID),
		zap.String(keyUsername, m.Username),
		zap.String(keyClientIP, m.ClientIP),
		zap.String(keyDeviceID, m.DeviceID),
	}
	for k, v := range m.ExtData {
		extraFields = append(extraFields, zap.Any(k, v))
	}
	return extraFields
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
	msg = &Message{Context: c, StartTime: time.Now(), ExtData: make(map[string]interface{}, 0)}
	c.Set(KeyMessage, msg)
	return msg
}

func MessageFromCtx(c context.Context) *Message {
	if ginCtx, ok := c.(*gin.Context); ok {
		return GinMessage(ginCtx)
	}
	return nil
}
