package database

import (
	"github.com/gallifreyCar/gallifrey-wechat/models"
	"gorm.io/gorm"
)

type UserDao struct{}

func NewUserDao() *UserDao { return &UserDao{} }

// IsEmailExist  邮箱是否存在
func (u *UserDao) IsEmailExist(email string) bool {
	var count int64
	DB.Model(models.UserBasic{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist IsPhoneExist 手机号是否存在
func (u *UserDao) IsPhoneExist(phone string) bool {
	var count int64
	DB.Model(models.UserBasic{}).Where("phone = ?", phone).Count(&count)
	return count > 0
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
	return DB.Where("identity = ?", user.Identity).Delete(&user)
}

// GetUser 获取用户
func (u *UserDao) GetUser(user *models.UserBasic) *gorm.DB {
	return DB.Where(user).First(&user)
}

// GetUsers 获取所有用户
func (u *UserDao) GetUsers(users *[]models.UserBasic) *gorm.DB {
	return DB.Find(&users)
}
