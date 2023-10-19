package database

import (
	"github.com/gallifreyCar/gallifrey-wechat/pkg/logger"
	"github.com/gallifreyCar/gallifrey-wechat/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() error {
	// 数据库配置
	cfg := utils.ReadMySQLConfig("config")

	// 连接数据库
	db, err := gorm.Open(mysql.Open(cfg), &gorm.Config{
		Logger: logger.NewGormLogger(logger.Logger, gormLogger.Info),
	})
	if err != nil {
		return err
	}

	DB = db
	return nil

}
