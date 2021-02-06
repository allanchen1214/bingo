package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DefaultLogger 默认的logger
var DefaultLogger *zap.Logger

func init() {
	DefaultLogger = RotateFileLogger(DefaultOption())
}

func log(logger *zap.Logger, level Level, msg string, fields ...zapcore.Field) {
	switch level {
	case debugLevel:
		logger.Debug(msg, fields...)
	case warnLevel:
		logger.Warn(msg, fields...)
	case errorLevel:
		logger.Error(msg, fields...)
	case fatalLevel:
		logger.Fatal(msg, fields...)
	default:
		logger.Info(msg, fields...)
	}
}

func Debug(msg string, fields ...zapcore.Field) {
	log(DefaultLogger, debugLevel, msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	log(DefaultLogger, infoLevel, msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	log(DefaultLogger, warnLevel, msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	log(DefaultLogger, errorLevel, msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	log(DefaultLogger, fatalLevel, msg, fields...)
}
