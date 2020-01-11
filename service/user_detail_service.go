package service

import (
	"island/model"
	"island/serializer"
)

// 获取用户详情
type UserDetailService struct{}

func (service *UserDetailService) Get(id string) (*model.User, *serializer.ErrorResponse) {
	var user model.User
	if err := model.DB.First(&user, id).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    4004,
			Message: "没有找到该用户",
			Error:   err.Error(),
		}
	}
	return &user, nil
}
