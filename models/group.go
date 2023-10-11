package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	OwnerID   uint   `json:"owner_id"`   // 拥有者ID
	GroupName string `json:"group_name"` // 群名称
	GroupDesc string `json:"group_desc"` // 群描述
	Icon      string `json:"icon"`       // 群头像
}
