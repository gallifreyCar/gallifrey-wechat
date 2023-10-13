package initialize

import "github.com/gallifreyCar/gallifrey-wechat/database"

func InitDB() {
	// 初始化数据库
	database.Init()
	//// 初始化表
	database.InitTable()
}
