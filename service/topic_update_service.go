package service

import (
	"island/model"
	"island/serializer"
)

type TopicUpdateService struct {
	Title string `json:"title" binding:"required"` // 标题
	Info  string `json:"info"`                     // 详情
	Icon  string `json:"icon" binding:"required"`  // 图标
	Image string `json:"image"`                    // 背景
}

func (service *TopicUpdateService) Update(id string, userId uint) (*model.Topic, *serializer.ErrorResponse) {

	var topic model.Topic
	if err := model.DB.First(&topic, id).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    4004,
			Message: "没有找到该话题",
			Error:   err.Error(),
		}
	}
	if topic.CreatorID != userId {
		return nil, &serializer.ErrorResponse{
			Code:    4003,
			Message: "不是该话题的创建者，无权修改",
		}
	}

	topic.Title = service.Title
	topic.Icon = service.Icon
	topic.Info = service.Info
	topic.Image = service.Image

	if err := model.DB.Save(&topic).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库保存出错",
			Error:   err.Error(),
		}
	}
	return &topic, nil
}
