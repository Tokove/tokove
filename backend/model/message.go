package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderID   uint   `json:"sender_id" gorm:"not null;index"`
	ReceiverID uint   `json:"receiver_id" gorm:"index"` // 可选单聊和群聊
	GroupID    uint   `json:"group_id" gorm:"index"`
	Content    string `json:"content" gorm:"type:text;not null"`
	IsRead     bool   `json:"is_read" gorm:"default:false"`
}
