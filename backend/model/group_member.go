package model

import "gorm.io/gorm"

type GroupMember struct {
	gorm.Model
	GroupID uint   `json:"group_id" gorm:"not null;index;uniqueIndex:idx_group_user"`
	UserID  uint   `json:"user_id" gorm:"not null;index;uniqueIndex:idx_group_user"`
	Role    string `json:"role" gorm:"size:32;default:'member'"` // 群主/管理员/普通成员
}
