// Package log 封装全局log
package log

import (
	"github.com/topfreegames/pitaya/v2/client"
	"github.com/topfreegames/pitaya/v2/logger"
	"go.uber.org/zap"
)

var Log = loggerMgr.Log     // 全局log
var Sugar = loggerMgr.Sugar // 全局Sugar

// loggerMgr 框架内部log持有管理类
var loggerMgr = logger.NewLogger(zap.NewProductionConfig(), logger.WithStackWithFmtFormatter(true))

func SetLevel(level any) {
	loggerMgr.SetLevel(level)
	client.SetLogLevel(level)
}
func SetDevelopment(enable bool) {
	loggerMgr.SetDevelopment(enable)
	Log = loggerMgr.Log
	Sugar = loggerMgr.Sugar
	client.SetLogDevelopment(enable)
}
