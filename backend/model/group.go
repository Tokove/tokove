package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name        string `json:"name" gorm:"size:128;not null;uniqueIndex"`
	Description string `json:"description" gorm:"size:255"`
	CreatorID   uint   `json:"creator_id" gorm:"not null;index"`
	Creator     User   `json:"creator" gorm:"foreignKey:CreatorID"`
	MemberCount int64  `json:"member_count" gorm:"default:0"`
}
