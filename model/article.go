package model

import (
	"github.com/jinzhu/gorm"
)

// Article 文章模型
type Article struct {
	gorm.Model
	Title     string `gorm:"size:100;not null;default:''"`
	Image     string `gorm:"size:100;not null;default:''"`
	Tags      string `gorm:"size:100;not null;default:''"`
	Profile   string `gorm:"size:100;not null;default:''"`
	Content   string `gorm:"type:longtext;not null'"`
	Topic     Topic  `gorm:"ForeignKey:TopicID"`
	Creator   User   `gorm:"ForeignKey:CreatorID"`
	TopicID   uint   `gorm:"not null;index"`
	CreatorID uint   `gorm:"not null;index"`
}
