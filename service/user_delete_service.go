package service

import (
	"island/model"
	"island/serializer"
)

type UserDeleteService struct{}

func (service *UserDeleteService) Delete(id uint) *serializer.ErrorResponse {
	var user model.User
	if err := model.DB.First(&user, id).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    4004,
			Message: "没有找到该用户",
			Error:   err.Error(),
		}
	}
	if err := model.DB.Delete(&user).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    5000,
			Message: "删除用户失败",
			Error:   err.Error(),
		}
	}
	return nil
}
