package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/Alpensin/go-obsmonster/pkg/logging/console"
)

const LoggerName = "logrus"

// logrusLogger wrapper, that will meet common interface for all loggers.
// It will translate our messages to format logger expect
type logrusLogger struct {
	logger *logrus.Logger
}

func New() console.Logger {
	logger := logrus.New()
	logger.AddHook(&loggerHook{})
	logger.SetFormatter(&logrus.JSONFormatter{})
	return &logrusLogger{
		logger: logger,
	}
}

func (l *logrusLogger) Debug(msg string, args ...console.Arg) {
	fields := l.convertArgs(args...)
	l.logger.WithFields(fields).Debug(msg)
}

func (l *logrusLogger) Info(msg string, args ...console.Arg) {
	fields := l.convertArgs(args...)
	l.logger.WithFields(fields).Info(msg)
}

func (l *logrusLogger) Warn(msg string, args ...console.Arg) {
	fields := l.convertArgs(args...)
	l.logger.WithFields(fields).Warn(msg)
}

func (l *logrusLogger) Error(msg string, args ...console.Arg) {
	fields := l.convertArgs(args...)
	l.logger.WithFields(fields).Error(msg)
}

func (l *logrusLogger) convertArgs(args ...console.Arg) logrus.Fields {
	attrs := make(map[string]any, len(args))
	for _, arg := range args {
		attrs[arg.Key] = arg.Value
	}
	return attrs
}

type loggerHook struct{}

func (h *loggerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *loggerHook) Fire(entry *logrus.Entry) error {
	entry.Data[console.LoggerNameField] = LoggerName
	return nil
}
