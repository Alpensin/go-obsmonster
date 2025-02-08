package slog

import (
	"context"
	"log/slog"

	"github.com/Alpensin/go-obsmonster/pkg/logging/stdout"
)

// slogLogger wrapper, that will meet common interface for all loggers.
// It will translate our messages to format logger expect
type slogLogger struct {
	logger slog.Logger
}

func New() stdout.Logger {
	return &slogLogger{}
}

func (l *slogLogger) Debug(msg string, args ...stdout.Arg) {
	attrs := l.convertArgs(args...)
	l.logger.LogAttrs(context.Background(), slog.LevelDebug, msg, attrs...)
}
func (l *slogLogger) Info(msg string, args ...stdout.Arg) {
	attrs := l.convertArgs(args...)
	l.logger.LogAttrs(context.Background(), slog.LevelInfo, msg, attrs...)
}
func (l *slogLogger) Warn(msg string, args ...stdout.Arg) {
	attrs := l.convertArgs(args...)
	l.logger.LogAttrs(context.Background(), slog.LevelWarn, msg, attrs...)
}
func (l *slogLogger) Critical(msg string, args ...stdout.Arg) {
	attrs := l.convertArgs(args...)
	l.logger.LogAttrs(context.Background(), slog.LevelError, msg, attrs...)
}

func (l *slogLogger) convertArgs(args ...stdout.Arg) []slog.Attr {
	attrs := make([]slog.Attr, 0, len(args))
	for _, arg := range args {
		attrs = append(attrs, slog.Any(arg.Key, arg.Value))
	}
	return attrs
}
