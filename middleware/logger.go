package middleware

import (
	"cblog/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

func LoggerToFile() gin.HandlerFunc {

	logFilePath := setting.LogSetting.Path
	_, err := os.Stat(logFilePath)
	if err != nil {
		err = os.MkdirAll(logFilePath, os.ModePerm)
		if err != nil {
			log.Printf("[fail] cannot make dir %s", logFilePath)
		}
	}
	logFileName := fmt.Sprintf("%s%s.%s", "log", time.Now().Format("20060102"), "log")
	// log file
	logFile := path.Join(logFilePath, logFileName)
	// write to log file
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		log.Printf("[fail] open logger file error: %v", err)
	}

	// log init
	logger := logrus.New()
	logger.Out = f
	logger.Level = logrus.DebugLevel
	logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方法
		requestMethod := c.Request.Method
		// 请求地址
		requestUrl := c.Request.RequestURI
		// 请求参数
		requestParams := c.Request.PostForm
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		requestIp := c.ClientIP()
		// 记录日志
		logger.WithFields(logrus.Fields{
			"code":         statusCode,
			"latency_time": latencyTime,
			"method":       requestMethod,
			"url":          requestUrl,
			"params":       requestParams,
			"ip":           requestIp,
		}).Info()
	}
}
