package service

import (
	"fmt"
	"island/model"
	"island/serializer"
)

type TopicCreateService struct {
	Title string `json:"title" binding:"required"` // 标题
	Info  string `json:"info"`                     // 详情
	Icon  string `json:"icon" binding:"required"`  // 图标
	Image string `json:"image"`                    // 背景
}

func (service *TopicCreateService) Create(id uint) (*model.Topic, *serializer.ErrorResponse) {
	topic := &model.Topic{
		Title:     service.Title,
		Info:      service.Info,
		Icon:      service.Icon,
		Image:     service.Image,
		CreatorID: id,
	}

	if err := model.DB.Create(&topic).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    4125,
			Message: "新建话题失败",
			Error:   fmt.Sprint(err),
		}
	}

	return topic, nil
}
