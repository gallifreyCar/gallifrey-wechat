package main

import (
	"github.com/gallifreyCar/gallifrey-wechat/models"
	"gorm.io/driver/mysql"
	"os"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	// 数据库配置
	cfg := mysqlCfg.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("MYSQL_ADDR"),
		DBName:               "gallifrey_wechat",
		ParseTime:            true,
		AllowNativePasswords: true,
		Loc:                  time.Local,
	}

	// 连接数据库
	db, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

}
