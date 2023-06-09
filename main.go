package main

import (
	"dbOper/routers"
	"dbOper/service"
	"dbOper/util"
	"embed"
)

//go:embed resources/*.sql
var f embed.FS

func main() {
	//连接数据库
	service.InitMySql()
	//程序退出关闭数据库连接
	defer service.Close()
	//初始化表结构
	service.CreateTable(f)

	//注册路由
	r := routers.SetRouter()
	//启动端口
	r.Run(util.ApplicationConfig.GetString("server.HttpPort"))
}
