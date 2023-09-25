package main

import (
	"fmt"
	"github.com/gallifreyCar/gallifrey-wechat/models"
	"github.com/gallifreyCar/gallifrey-wechat/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 数据库配置
	cfg := utils.ReadMySQLConfig("config")
	fmt.Println(cfg)
	// 连接数据库
	db, err := gorm.Open(mysql.Open(cfg), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

}
