package service

import (
	"island/model"
	"island/serializer"
)

type TopicDeleteService struct{}

func (service *TopicDeleteService) Delete(id string, userId uint) *serializer.ErrorResponse {
	var topic model.Topic
	if err := model.DB.First(&topic, id).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    4004,
			Message: "没有找到该话题",
			Error:   err.Error(),
		}
	}
	if topic.CreatorID != userId {
		return &serializer.ErrorResponse{
			Code:    4003,
			Message: "不是该话题的创建者，无权删除",
		}
	}
	if err := model.DB.Delete(&topic).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    5000,
			Message: "删除失败",
			Error:   err.Error(),
		}
	}
	return nil
}
