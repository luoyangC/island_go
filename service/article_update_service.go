package service

import (
	"island/model"
	"island/serializer"
)

type ArticleUpdateService struct {
	Title   string `json:"title" binding:"required"`              // 标题
	Image   string `json:"image"`                                 // banner图
	Tags    string `json:"tags" binding:"required"`               // 标签
	Profile string `json:"profile" binding:"required"`            // 简介
	Content string `json:"content" binding:"required"`            // 内容
	TopicID uint   `json:"topicId" binding:"required.TopicValid"` // 话题ID
}

func (service *ArticleUpdateService) Update(id string, userId uint) (*model.Article, *serializer.ErrorResponse) {

	var article model.Article
	if err := model.DB.First(&article, id).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    4004,
			Message: "没有找到该文章",
			Error:   err.Error(),
		}
	}
	if article.CreatorID != userId {
		return nil, &serializer.ErrorResponse{
			Code:    4003,
			Message: "不是该文章的作者，无权修改",
		}
	}

	article.Title = service.Title
	article.Image = service.Image
	article.Tags = service.Tags
	article.Profile = service.Profile
	article.Content = service.Content
	article.TopicID = service.TopicID

	if err := model.DB.Save(&article).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库保存出错",
			Error:   err.Error(),
		}
	}
	return &article, nil
}
