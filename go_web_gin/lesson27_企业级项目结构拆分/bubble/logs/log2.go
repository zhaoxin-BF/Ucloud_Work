package logs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func InitLog(){
	//日志写入文件时，禁用控制台颜色
	gin.DisableConsoleColor()

	//写入日志文件
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	//如果需要同时写入日志文件和控制台上显示，使用下面代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	fmt.Fprintln(gin.DefaultWriter, "--------------------foo bar----------------")
}

func Info(str string){
	opttime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintln(gin.DefaultWriter, "[INFO]    |"+opttime+"|   ------------Info |:  "+str)
}

func Error(str string){
	opttime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintln(gin.DefaultWriter, "[Error]   |"+opttime+"|   ++++++++++++Error|:  "+str)
}