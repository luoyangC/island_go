package service

import (
	"github.com/gin-gonic/gin"

	"island/model"
	"island/serializer"
	"island/utils"
)

type UserListService struct{}

// List 获取用户列表
func (service *UserListService) List(c *gin.Context) ([]*model.User, int, *serializer.ErrorResponse) {
	var users []*model.User
	var count int

	limit, offset, err := utils.Pagination(c)
	if err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    4000,
			Message: "错误的分页",
			Error:   err.Error(),
		}
	}

	if err := model.DB.Model(model.User{}).Count(&count).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	if err := model.DB.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	return users, count, nil
}
