package logs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func Logger() gin.HandlerFunc {
	logClient := log.New()
	var logPath = "./"// 日志打印到指定的目录
	// 目录不存在则创建
	//if !util.PathExists(logPath) {
	//	os.MkdirAll(logPath, os.ModePerm)
	//}
	fileName := path.Join(logPath, "mikasa_ufs.log")
	//禁止logrus的输出
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err!= nil{
		fmt.Println("err", err)
	}
	// 设置日志输出的路径
	logClient.Out = src
	logClient.SetLevel(log.DebugLevel)
	//apiLogPath := "gin-api.log"
	logWriter, err := rotatelogs.New(
		"mikasa_ufs-%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(fileName), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &log.TextFormatter{})   //json格式打印日志
	logClient.AddHook(lfHook)

	//logClient.SetFormatter(&log.TextFormatter{})              //文本格式打印日志


	return func (c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		// 这里是指定日志打印出来的格式。分别是状态码，执行时间,请求ip,请求方法,请求路由(等下我会截图)
		logClient.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}