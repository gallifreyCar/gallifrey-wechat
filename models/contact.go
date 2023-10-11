package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	OwnerID      uint `json:"owner_id"`     // 拥有者ID
	ContactID    uint `json:"contact_id"`   // 联系人ID
	Relationship int  `json:"relationship"` // 0: 好友 1: 关注 2: 粉丝 3: 黑名单
}
