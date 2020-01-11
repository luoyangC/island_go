package service

import (
"island/model"
"island/serializer"
)

type TopicDetailService struct{}

func (service *TopicDetailService) Get(id string) (*model.Topic, *serializer.ErrorResponse) {
	var topic model.Topic
	if err := model.DB.First(&topic, id).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    4004,
			Message: "没有找到该话题",
			Error:   err.Error(),
		}
	}
	return &topic, nil
}
