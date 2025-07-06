package logger

import (
	"go.uber.org/zap"
)

func NewLogger(lg zap.Logger) *zap.Logger {
	domainLogger := lg.With(
		zap.String("service", "userService"),
		zap.String("requestID", "abc123"),
	)
	return domainLogger
}
