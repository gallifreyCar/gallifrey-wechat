package main

import (
	"github.com/gallifreyCar/gallifrey-wechat/database"
	"github.com/gallifreyCar/gallifrey-wechat/http/controller/v1"
	"github.com/gallifreyCar/gallifrey-wechat/initialize"
	"github.com/gallifreyCar/gallifrey-wechat/redis"
	"github.com/gallifreyCar/gallifrey-wechat/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	database.Init()
	//// 初始化表
	database.InitTable()
	// 初始化控制器
	v1.Init()
	// 初始化验证器
	service.InitValidator()
	// 初始化Redis
	redis.InitRedis()

	// 启动路由
	r := gin.New()
	initialize.InitRoute(r)
	r.Run()
}
