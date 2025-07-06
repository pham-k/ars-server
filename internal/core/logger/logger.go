// Credit https://github.com/dusted-go/logging/blob/main/prettylog/prettylog.go

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitializeGlobalLogger() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableCaller = false

	logger := zap.Must(config.Build())
	defer logger.Sync()
	zap.ReplaceGlobals(zap.Must(config.Build()))
}

func NewLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableCaller = false

	logger := zap.Must(config.Build())
	defer logger.Sync()

	return logger
}
