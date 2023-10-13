package database

import "github.com/gallifreyCar/gallifrey-wechat/models"

type AuthDao struct{}

// IsEmailExist  邮箱是否存在
func (a *AuthDao) IsEmailExist(email string) bool {
	var count int64
	DB.Model(models.UserBasic{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist IsPhoneExist 手机号是否存在
func (a *AuthDao) IsPhoneExist(phone string) bool {
	var count int64
	DB.Model(models.UserBasic{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
