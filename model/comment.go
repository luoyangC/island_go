package model

import (
	"github.com/jinzhu/gorm"
)

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content   string  `gorm:"not null;default:''"`
	Article   Article `gorm:"ForeignKey:ArticleID"`
	Creator   User    `gorm:"ForeignKey:CreatorID"`
	ArticleID uint    `gorm:"not null;index"`
	CreatorID uint    `gorm:"not null;index"`
}
