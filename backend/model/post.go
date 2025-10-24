package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID   uint   `json:"user_id" gorm:"not null;index"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`
	Topic    string `json:"topic" gorm:"size:64"`
	Content  string `json:"content" gorm:"type:text;not null"`
	Images   string `json:"images" gorm:"type:json"`
	Video    string `json:"video" gorm:"size:512"`
	Likes    int    `json:"likes" gorm:"default:0"`
	Comments int    `json:"comments" gorm:"default:0"`
	Views    int    `json:"views" gorm:"default:0"`
	Visible  string `json:"visible" gorm:"size:20;default:'public'"` // 可见范围: public/friends/private
	Location string `json:"location" gorm:"size:128"`
}
