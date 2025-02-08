package slog

import (
	"context"
	"log/slog"
	"os"

	"github.com/Alpensin/go-obsmonster/pkg/logging/console"
)

// slogLogger wrapper, that will meet common interface for all loggers.
// It will translate our messages to format logger expect
type slogLogger struct {
	logger *slog.Logger
}

func New() console.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	return &slogLogger{
		logger: logger,
	}
}

func (l *slogLogger) Debug(msg string, args ...console.Arg) {
	attrs := l.convertArgs(args...)
	l.logger.LogAttrs(context.Background(), slog.LevelDebug, msg, attrs...)
}
func (l *slogLogger) Info(msg string, args ...console.Arg) {
	attrs := l.convertArgs(args...)
	l.logger.LogAttrs(context.Background(), slog.LevelInfo, msg, attrs...)
}
func (l *slogLogger) Warn(msg string, args ...console.Arg) {
	attrs := l.convertArgs(args...)
	l.logger.LogAttrs(context.Background(), slog.LevelWarn, msg, attrs...)
}
func (l *slogLogger) Critical(msg string, args ...console.Arg) {
	attrs := l.convertArgs(args...)
	l.logger.LogAttrs(context.Background(), slog.LevelError, msg, attrs...)
}

func (l *slogLogger) convertArgs(args ...console.Arg) []slog.Attr {
	attrs := make([]slog.Attr, 0, len(args))
	for _, arg := range args {
		attrs = append(attrs, slog.Any(arg.Key, arg.Value))
	}
	return attrs
}
