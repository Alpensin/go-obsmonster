// Package mux is multiplexer for all loggers. Need to log at all loggers at once
package mux

import (
	"github.com/Alpensin/go-obsmonster/pkg/logging/console"
	"github.com/Alpensin/go-obsmonster/pkg/logging/console/logrus"
	"github.com/Alpensin/go-obsmonster/pkg/logging/console/slog"
	"github.com/Alpensin/go-obsmonster/pkg/logging/console/zap"
	"github.com/Alpensin/go-obsmonster/pkg/logging/console/zerolog"
)

// LoggersMux - multiplexer for all loggers. Helps to log same message to every logger at once
type LoggersMux []console.Logger

// New LoggersMux
func New() LoggersMux {
	return LoggersMux{
		slog.New(),
		zap.New(),
		zerolog.New(),
		logrus.New(),
	}
}

func (mx LoggersMux) Debug(msg string, args ...console.Arg) {
	for _, logger := range mx {
		logger.Debug(msg, args...)
	}
}

func (mx LoggersMux) Info(msg string, args ...console.Arg) {
	for _, logger := range mx {
		logger.Info(msg, args...)
	}
}

func (mx LoggersMux) Warn(msg string, args ...console.Arg) {
	for _, logger := range mx {
		logger.Warn(msg, args...)
	}
}

func (mx LoggersMux) Critical(msg string, args ...console.Arg) {
	for _, logger := range mx {
		logger.Critical(msg, args...)
	}
}
