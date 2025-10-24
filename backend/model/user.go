package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:64;uniqueIndex;not null"`
	Password string `json:"password" gorm:"size:255;not null"`
	Age      uint   `json:"age"`
	Address  string `json:"address"`
	Email    string `json:"email" gorm:"size:128;uniqueIndex"`
	Avatar   string `json:"avatar" gorm:"size:255;not null;default:'/static/avatar/avatar.png'"`
	Bio      string `json:"bio" gorm:"size:255"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}
