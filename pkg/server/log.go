package server

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func init() {
	newLogger, _ := zap.NewProduction()
	logger = newLogger.Sugar()
}
