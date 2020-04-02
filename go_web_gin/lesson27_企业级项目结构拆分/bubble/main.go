package main

import (
	"bubble/dao"
	"bubble/logs"
	"bubble/models"
	"bubble/routers"
)

func main() {
	logs.InitLog()
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 	1、连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()  // 程序退出关闭数据库连接

	// 2、模型绑定，绑定库db1中，表todo,没有则会创建
	dao.DB.AutoMigrate(&models.Todo{})


	// 4、注册路由
	r := routers.SetupRouter()
	r.Run()
}
