package database

import "github.com/gallifreyCar/gallifrey-wechat/models"

// InitTable 表不存在时创建表
func InitTable() error {
	// 迁移 schema
	return DB.AutoMigrate(&models.UserBasic{})
}
