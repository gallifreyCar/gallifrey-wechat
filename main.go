package main

import (
	"github.com/gallifreyCar/gallifrey-wechat/controller"
	"github.com/gallifreyCar/gallifrey-wechat/database"
	"github.com/gallifreyCar/gallifrey-wechat/router"
)

func main() {
	// 初始化数据库
	database.Init()
	// 初始化控制器
	controller.Init()
	r := router.Router()
	r.Run()
}
