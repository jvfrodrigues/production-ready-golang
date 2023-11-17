package zap

import (
	"go.uber.org/zap"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (lo Logger) Info(message string, args ...interface{}) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	if len(args) > 0 {
		sugar.Info(message, args)
	} else {
		sugar.Info(message)
	}
}

func (lo Logger) Error(message string, args ...interface{}) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	if len(args) > 0 {
		sugar.Error(message, args)
	} else {
		sugar.Error(message)
	}
}
