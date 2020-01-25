package model

import (
	"github.com/jinzhu/gorm"
)

// Reply 回复模型
type Reply struct {
	gorm.Model
	Content    string  `gorm:"not null;default:''"`
	Comment    Comment `gorm:"ForeignKey:CommentID"`
	Creator    User    `gorm:"ForeignKey:CreatorID"`
	Receiver   User    `gorm:"ForeignKey:ReceiverID"`
	CommentID  uint    `gorm:"not null;index"`
	CreatorID  uint    `gorm:"not null;index"`
	ReceiverID uint    `gorm:"not null;index"`
}
