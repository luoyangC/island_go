package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName string `gorm:"size:20;not null"`
	Password string `gorm:"not null"`
	Status   int    `gorm:"not null;default:1"`
	Mobile   string `gorm:"size:11;not null"`
	Email    string `gorm:"size:100;default:''"`
	Avatar   string `gorm:"size:200;default:''"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active int = 1
	// Suspend 被封禁用户
	Suspend int = 2
	// Destroyed 注销用户
	Destroyed int = 0
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
