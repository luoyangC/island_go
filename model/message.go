package model

import (
	"github.com/jinzhu/gorm"
)

// Message 留言模型
type Message struct {
	gorm.Model
	Content   string `gorm:"not null;default:''"`
	Email     string `gorm:"size:100;default:''"`
	Mobile    string `gorm:"size:11;default:''"`
	Avatar    string `gorm:"size:100;default:''"`
	NickName  string `gorm:"size:20;not null"`
	ClientIP  string `gorm:"size:100"`
}
