package model

import (
	"github.com/jinzhu/gorm"
)

// Collection 收藏模型
type Collection struct {
	gorm.Model
	Article   Article `gorm:"ForeignKey:ArticleID"`
	Creator   User    `gorm:"ForeignKey:CreatorID"`
	ArticleID uint    `gorm:"not null;index"`
	CreatorID uint    `gorm:"not null;index"`
}
