package loggersystem

import "go.uber.org/zap"

type Logger struct {
	*zap.SugaredLogger
}

// база логгера
func Init() *zap.SugaredLogger {
	BaseLogger, _ := zap.NewDevelopment()

	return BaseLogger.Sugar()
}
