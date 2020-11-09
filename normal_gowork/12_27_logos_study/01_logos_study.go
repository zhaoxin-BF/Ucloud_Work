package main

//访问配置文件（conf .json文件）

import (
	"time"

	"github.com/astaxie/beego/logs"
)

func main() {
	//输出到控制台
	logs.SetLogger("console")

	//输出文件名和文件行号，log.go的文件名和代码行号
	// logs.EnableFuncCallDepth(true)

	//设置异步的输出
	logs.Async(1e3)

	//输出到日志文档
	logs.SetLogger(logs.AdapterFile,
		`{"filename":"project.log",
		"level":7,
		"maxlines":0,
		"maxsize":0,
		"daily":true,
		"maxdays":10,
		"color":true}`)
	//一秒钟时间间隔打印logs.Info,logs.Error,logs.Worn,logs.
	timer := time.NewTicker(1 * time.Second) //循环定时器

	for i := 0; i < 5000; i++ {
		<-timer.C                                       //循环从channel中取数据
		t := time.Now().Format("2006年01月02日 15时04分05秒") //月日分秒的格式化必须是04零开头，不然就会出错
		logs.Info("查询到时间信息：", t)
		logs.Error("查询到错误时间：", t)
		logs.Warn("查询到警告时间：", t)
		logs.Debug("查询到调试时间：", t)
	}

}
