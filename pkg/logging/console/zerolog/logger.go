package zerolog

import (
	"os"

	"github.com/rs/zerolog"

	"github.com/Alpensin/go-obsmonster/pkg/logging/console"
)

const LoggerName = "zerolog"

// zerologLogger wrapper, that will meet common interface for all loggers.
// It will translate our messages to format logger expect
type zerologLogger struct {
	logger *zerolog.Logger
}

func New() console.Logger {
	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Str(console.LoggerNameField, LoggerName).
		Logger()
	return &zerologLogger{
		logger: &logger,
	}
}

func (l *zerologLogger) Debug(msg string, args ...console.Arg) {
	event := l.logger.Debug()
	for _, arg := range args {
		event.Any(arg.Key, arg.Value)
	}
	event.Msg(msg)
}

func (l *zerologLogger) Info(msg string, args ...console.Arg) {
	event := l.logger.Info()
	for _, arg := range args {
		event.Any(arg.Key, arg.Value)
	}
	event.Msg(msg)
}

func (l *zerologLogger) Warn(msg string, args ...console.Arg) {
	event := l.logger.Warn()
	for _, arg := range args {
		event.Any(arg.Key, arg.Value)
	}
	event.Msg(msg)
}

func (l *zerologLogger) Error(msg string, args ...console.Arg) {
	event := l.logger.Error()
	for _, arg := range args {
		event.Any(arg.Key, arg.Value)
	}
	event.Msg(msg)
}
