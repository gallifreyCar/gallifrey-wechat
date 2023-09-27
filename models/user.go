package models

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string `validate:"required"` //校验用户名非空
	Password      string `validate:"required"` //校验密码
	Phone         string //`validate:"max=11,min=11,numeric"` //校验手机号
	Email         string //`validate:"email"`                 //校验邮箱号
	Identity      string
	ClientIP      string //`validate:"ip"`
	ClientPort    string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LogoutTime    *time.Time
	isLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
