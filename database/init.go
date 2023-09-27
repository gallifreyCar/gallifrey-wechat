package database

import (
	"fmt"
	"github.com/gallifreyCar/gallifrey-wechat/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func Init() error {
	// 数据库配置
	cfg := utils.ReadMySQLConfig("config")
	fmt.Println(cfg)

	// 日志配置
	mysqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			//慢查询阈值
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})

	// 连接数据库
	db, err := gorm.Open(mysql.Open(cfg), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
