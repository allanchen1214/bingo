package middleware

import (
	"time"

	"github.com/allanchen1214/bingo/log"
	"github.com/allanchen1214/bingo/trace"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	// KeyBenchmark bm
	KeyBenchmark = "Benchmark"

	keyHandlerName = "HandlerName"
	keyMethod      = "Method"
	keyFullPath    = "FullPath"
	keyPreLatency  = "PreLatency"
	keyRecvTime    = "RecvTime"
	keyLatency     = "Latency"
	keyResult      = "Result"
)

// Benchmark 记录处理耗时以及结果
func Benchmark() gin.HandlerFunc {
	return func(c *gin.Context) {
		msg := trace.GinMessage(c)
		recvTime := time.Now()
		preLatency := recvTime.Sub(msg.StartTime)

		c.Next()

		latency := time.Since(recvTime)
		log.InfoWithContext(c, KeyBenchmark, zap.String(trace.KeySequence, msg.Sequence),
			zap.String(trace.KeyUserID, msg.UserID),
			zap.String(trace.KeyUsername, msg.Username),
			zap.String(trace.KeyClientIP, msg.ClientIP),
			zap.String(trace.KeyDeviceID, msg.DeviceID),
			zap.String(keyHandlerName, c.HandlerName()),
			zap.String(keyMethod, c.Request.Method),
			zap.String(keyFullPath, c.FullPath()),
			zap.Time(trace.KeyStartTime, msg.StartTime),         // 后端首次接收到请求时间
			zap.Int64(keyPreLatency, preLatency.Milliseconds()), // 进入到本环节之前所有环节处理耗时
			zap.Time(keyRecvTime, recvTime),
			zap.Int64(keyLatency, latency.Milliseconds()), // 本环节处理耗时
			zap.Int(keyResult, c.Writer.Status()),
		)
	}
}
