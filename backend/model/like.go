package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID   uint   `json:"user_id" gorm:"not null;index;uniqueIndex:idx_user_target_type"` // 组合唯一索引，防止多次点赞
	TargetID uint   `json:"target_id" gorm:"not null;index;uniqueIndex:idx_user_target_type"`
	Type     string `json:"type" gorm:"size:32;not null;uniqueIndex:idx_user_target_type"`
}
