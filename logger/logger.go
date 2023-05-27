package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger() {
	Logger = zap.L()
}
