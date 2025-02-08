package zap

import (
	"go.uber.org/zap"

	"github.com/Alpensin/go-obsmonster/pkg/logging/console"
)

// zapLogger wrapper, that will meet common interface for all loggers.
// It will translate our messages to format logger expect
type zapLogger struct {
	logger *zap.Logger
}

func New() console.Logger {
	logger, _ := zap.NewProduction()
	return &zapLogger{
		logger: logger,
	}
}

func (l *zapLogger) Debug(msg string, args ...console.Arg) {
	fields := l.convertArgs(args...)
	l.logger.Debug(msg, fields...)
}
func (l *zapLogger) Info(msg string, args ...console.Arg) {
	fields := l.convertArgs(args...)
	l.logger.Info(msg, fields...)
}
func (l *zapLogger) Warn(msg string, args ...console.Arg) {
	fields := l.convertArgs(args...)
	l.logger.Warn(msg, fields...)
}
func (l *zapLogger) Critical(msg string, args ...console.Arg) {
	fields := l.convertArgs(args...)
	l.logger.Error(msg, fields...)
}

func (l *zapLogger) convertArgs(args ...console.Arg) []zap.Field {
	attrs := make([]zap.Field, 0, len(args))
	for _, arg := range args {
		attrs = append(attrs, zap.Any(arg.Key, arg.Value))
	}
	return attrs
}
