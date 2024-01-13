package log

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

// prevent logger initialization every request
func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Info(message string) {
	l.logger.Info(message)
}

func (l *Logger) Error(message string) {
	l.logger.Error(message)
}
