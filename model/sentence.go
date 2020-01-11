package model

import (
	"github.com/jinzhu/gorm"
)

// Sentence 句子模型
type Sentence struct {
	gorm.Model
	Lines     string `gorm:"size:200;not null;default:''"`
	Creator   User   `gorm:"ForeignKey:CreatorID"`
	CreatorID uint   `gorm:"not null;index"`
}
