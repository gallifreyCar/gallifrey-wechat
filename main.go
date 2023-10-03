package main

import (
	"github.com/gallifreyCar/gallifrey-wechat/controller"
	"github.com/gallifreyCar/gallifrey-wechat/database"
	"github.com/gallifreyCar/gallifrey-wechat/redis"
	"github.com/gallifreyCar/gallifrey-wechat/router"
	"github.com/gallifreyCar/gallifrey-wechat/service"
)

func main() {
	// 初始化数据库
	database.Init()
	//// 初始化表
	//database.InitTable()
	// 初始化控制器
	controller.Init()
	// 初始化验证器
	service.InitValidator()
	// 初始化Redis
	redis.InitRedis()

	r := router.Router()
	r.Run()
}
