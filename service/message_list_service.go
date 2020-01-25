package service

import (
	"github.com/gin-gonic/gin"

	"island/model"
	"island/serializer"
	"island/utils"
)

type MessageListService struct{}

func (service *MessageListService) List(c *gin.Context) ([]*serializer.Message, int, *serializer.ErrorResponse) {
	var count int
	limit, offset, err := utils.Pagination(c)
	if err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    4000,
			Message: "错误的分页",
			Error:   err.Error(),
		}
	}
	if err := model.DB.Model(model.Message{}).Count(&count).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	var messages []*serializer.Message
	if err := model.DB.Limit(limit).Offset(offset).Find(&messages).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库错误",
			Error:   err.Error(),
		}
	}
	return messages, count, nil
}
