package service

import (
	"island/model"
	"island/serializer"
)

type UserUpdateService struct {
	UserName string `json:"username"` // 用户名
	Email    string `json:"email"`    // 邮箱
	Avatar   string `json:"avatar"`   // 头像
}

func (service *UserUpdateService) Update(id uint) (*model.User, *serializer.ErrorResponse) {

	var user model.User
	if err := model.DB.First(&user, id).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    4004,
			Message: "没有找到该用户",
			Error:   err.Error(),
		}
	}

	user.UserName = service.UserName
	user.Email = service.Email
	user.Avatar = service.Avatar

	if err := model.DB.Save(&user).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库保存出错",
			Error:   err.Error(),
		}
	}
	return &user, nil
}
