package model

import (
	"github.com/jinzhu/gorm"
)

// Topic 话题模型
type Topic struct {
	gorm.Model
	Title     string `gorm:"size:20;not null;default:''"`
	Info      string `gorm:"size:100;not null;default:''"`
	Icon      string `gorm:"size:100;default:''"`
	Image     string `gorm:"size:100;default:''"`
	Creator   User   `gorm:"ForeignKey:CreatorID"`
	CreatorID uint   `gorm:"not null;index"`
}
