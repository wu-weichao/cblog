package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func init() {
	// 设置日志格式 默认为 log.TextFormatter{}
	Logger.Formatter = &logrus.JSONFormatter{}
	// 设置日志默认输出 默认为 os.Stderr 标准错误输出
	Logger.Out = os.Stdout
	// 设置日志记录级别
	Logger.Level = logrus.TraceLevel

	//// 设置日志格式 默认为 log.TextFormatter{}
	//log.SetFormatter(&log.JSONFormatter{})
	//// 设置日志默认输出 默认为 os.Stderr 标准错误输出
	//log.SetOutput(os.Stdout)
	//// 设置日志记录级别
	//log.SetLevel(log.TraceLevel)
}
