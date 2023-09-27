package database

import (
	"github.com/gallifreyCar/gallifrey-wechat/models"
	"gorm.io/gorm"
)

type UserDao struct{}

func NewUserDao() *UserDao { return &UserDao{} }

// CreateTable 表不存在时创建表
func (u *UserDao) CreateTable() error {
	// 迁移 schema
	return DB.AutoMigrate(&models.UserBasic{})
}

// CreateUser 插入用户
func (u *UserDao) CreateUser(user *models.UserBasic) *gorm.DB {
	return DB.Create(user)
}

// UpdateUser 更新用户
func (u *UserDao) UpdateUser(user *models.UserBasic) *gorm.DB {
	return DB.Model(&user).Updates(user)
}

// DeleteUser 删除用户
func (u *UserDao) DeleteUser(user *models.UserBasic) *gorm.DB {
	return DB.Where("name = ?", user.Name).Delete(&user)
}

// GetUser 获取用户
func (u *UserDao) GetUser(user *models.UserBasic) *gorm.DB {
	return DB.Where(user).First(&user)
}

// GetUsers 获取所有用户
func (u *UserDao) GetUsers(users *[]models.UserBasic) *gorm.DB {
	return DB.Find(&users)
}
