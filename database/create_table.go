package database

import "github.com/gallifreyCar/gallifrey-wechat/models"

// InitTable 表不存在时创建表
func InitTable() error {
	// 迁移 schema
	var err error
	err = DB.AutoMigrate(&models.UserBasic{})
	err = DB.AutoMigrate(&models.Group{})
	err = DB.AutoMigrate(&models.Contact{})
	if err != nil {
		return err
	}
	return nil
}
