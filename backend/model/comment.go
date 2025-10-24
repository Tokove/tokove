package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"not null;index"`
	User      User   `json:"user" gorm:"foreignKey:UserID"`
	PostID    uint   `json:"post_id" gorm:"not null;index"`
	Content   string `json:"content" gorm:"type:text;not null"`
	ParentID  *uint  `json:"parent_id" gorm:"index"` // 回复评论（nil表示一级评论）
	LikeCount int64  `json:"like_count" gorm:"default:0"`
}
