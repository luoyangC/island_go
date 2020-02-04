package service

import (
	"island/model"
	"island/serializer"
)

type TopicListService struct {}

func (service *TopicListService) List(limit int, offset int) ([]*model.Topic, int, *serializer.ErrorResponse) {
	var topics []*model.Topic
	var count int
	if err := model.DB.Model(model.Topic{}).Count(&count).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	if err := model.DB.Offset(offset).Limit(limit).Find(&topics).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	return topics, count, nil
}
