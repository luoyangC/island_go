package service

import (
	"island/model"
	"island/serializer"
)

type ArticleCreateService struct {
	Title   string `json:"title" binding:"required"`   // 标题
	Image   string `json:"image"`                      // banner图
	Tags    string `json:"tags" binding:"required"`    // 标签
	Profile string `json:"profile" binding:"required"` // 简介
	Content string `json:"content" binding:"required"` // 内容
	TopicID uint   `json:"topicId" binding:"required"` // 话题ID
}

func (service *ArticleCreateService) Create(userId uint) (*model.Article, *serializer.ErrorResponse) {
	article := model.Article{
		Title:     service.Title,
		Image:     service.Image,
		Tags:      service.Tags,
		Profile:   service.Profile,
		Content:   service.Content,
		TopicID:   service.TopicID,
		CreatorID: userId,
	}

	if err := model.DB.Create(&article).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    5001,
			Message: "创建文章失败",
			Error:   err.Error(),
		}
	}

	return &article, nil
}
