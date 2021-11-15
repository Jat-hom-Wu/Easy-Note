package main

import (
	"easy_note/router"
	"easy_note/model"
	"easy_note/dao"
	"fmt"
)

var port int = 8000
var setting dao.MysqlSetting

func HandleSetting(){
	setting.User = "root"
	setting.Password = "12345"
	setting.Host = "127.0.0.1"
	setting.Port = "3306"
	setting.DBname = "easy_note"
}

func main(){
	//解析命令行数据

	//初始化配置
	HandleSetting()
	//数据库连接
	db,_ := dao.InitMysql(setting)
	model.DB = db
	//路由处理
	routers := router.RoutersInit()
	if err := routers.Run(fmt.Sprintf(":%d", port)); err != nil{
		fmt.Println("routers run failed")
	}
}