package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableCaller = false

	logger := zap.Must(config.Build())

	defer logger.Sync()

	logger.Info("Hello from zap logger")
	return logger
}
