package models

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      string
	ClientIP      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogoutTime    time.Time
	isLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
