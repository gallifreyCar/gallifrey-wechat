package database

import (
	"fmt"
	"github.com/gallifreyCar/gallifrey-wechat/models"
	"github.com/gallifreyCar/gallifrey-wechat/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitDB() error {
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

// CreateTable 表不存在时创建表
func CreateTable() error {
	// 迁移 schema
	return DB.AutoMigrate(&models.UserBasic{})
}

// CreateUser 插入用户
func CreateUser(user *models.UserBasic) *gorm.DB {
	return DB.Create(user)
}

// UpdateUser 更新用户
func UpdateUser(user *models.UserBasic) *gorm.DB {
	return DB.Model(&user).Updates(user)
}

// DeleteUser 删除用户
func DeleteUser(user *models.UserBasic) *gorm.DB {
	return DB.Delete(&user)
}

// GetUser 获取用户
func GetUser(user *models.UserBasic) *gorm.DB {
	return DB.Table("user_basic").First(&user)
}

// GetUsers 获取所有用户
func GetUsers(users *[]models.UserBasic) *gorm.DB {
	return DB.Table("user_basic").Find(&users)
}