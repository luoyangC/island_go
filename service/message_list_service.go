package service

import (
	"island/model"
	"island/serializer"
)

type MessageListService struct{}

func (service *MessageListService) List(limit int, offset int) ([]*serializer.Message, int, *serializer.ErrorResponse) {
	var count int
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
