package database

import (
	"github.com/gallifreyCar/gallifrey-wechat/models"
	"gorm.io/gorm"
)

type UserDB interface {
	// CreateTable 表不存在时创建表
	CreateTable() error

	// CreateUser 插入用户
	CreateUser(user *models.UserBasic) *gorm.DB

	// UpdateUser 更新用户
	UpdateUser(user *models.UserBasic) *gorm.DB
	// DeleteUser 删除用户
	DeleteUser(user *models.UserBasic) *gorm.DB

	// GetUser 获取用户
	GetUser(user *models.UserBasic) *gorm.DB

	// GetUsers 获取所有用户
	GetUsers(users *[]models.UserBasic) *gorm.DB
}
