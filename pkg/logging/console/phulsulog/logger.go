package phulsulog

import (
	"github.com/phuslu/log"

	"github.com/Alpensin/go-obsmonster/pkg/logging/console"
)

const LoggerName = "phulsulog"

// phulsuLogger wrapper, that will meet common interface for all loggers.
// It will translate our messages to format logger expect
type phulsuLogger struct {
	logger *log.Logger
}

func New() console.Logger {
	logger := &log.Logger{
		Level:   log.DebugLevel,
		Context: log.NewContext(nil).Str(console.LoggerNameField, LoggerName).Value(),
	}
	return &phulsuLogger{
		logger: logger,
	}
}

func (l *phulsuLogger) Debug(msg string, args ...console.Arg) {
	entry := l.logger.Debug()
	for _, arg := range args {
		entry.Any(arg.Key, arg.Value)
	}
	entry.Msg(msg)
}
func (l *phulsuLogger) Info(msg string, args ...console.Arg) {
	entry := l.logger.Info()
	for _, arg := range args {
		entry.Any(arg.Key, arg.Value)
	}
	entry.Msg(msg)
}
func (l *phulsuLogger) Warn(msg string, args ...console.Arg) {
	entry := l.logger.Warn()
	for _, arg := range args {
		entry.Any(arg.Key, arg.Value)
	}
	entry.Msg(msg)
}
func (l *phulsuLogger) Error(msg string, args ...console.Arg) {
	entry := l.logger.Error()
	for _, arg := range args {
		entry.Any(arg.Key, arg.Value)
	}
	entry.Msg(msg)
}
