package model

import (
	"github.com/jinzhu/gorm"
)

// Like 点赞模型
type Like struct {
	gorm.Model
	LikeType  string `gorm:"size:10"`
	LikeID    uint   `gorm:"not null"`
	Creator   User   `gorm:"ForeignKey:CreatorID"`
	CreatorID uint   `gorm:"not null;index"`
}
