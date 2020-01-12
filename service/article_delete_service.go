package service

import (
	"island/model"
	"island/serializer"
)

type ArticleDeleteService struct{}

func (service *ArticleDeleteService) Delete(id string, userId uint) *serializer.ErrorResponse {
	var article model.Article
	if err := model.DB.First(&article, id).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    4004,
			Message: "没有找到该文章",
			Error:   err.Error(),
		}
	}
	if article.CreatorID != userId {
		return &serializer.ErrorResponse{
			Code:    4003,
			Message: "不是该文章的作者，无权删除",
		}
	}
	if err := model.DB.Delete(&article).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    5001,
			Message: "删除失败",
			Error:   err.Error(),
		}
	}
	return nil
}
