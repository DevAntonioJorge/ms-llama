package logger

import (
	"go.uber.org/zap"
)


func NewLogger(level string) *zap.Logger {
	config := zap.NewProductionConfig()
	if level == "debug" {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("Logger initialized", zap.String("level", level))	
	return logger
}