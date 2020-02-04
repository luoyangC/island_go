package service

import (
	"island/model"
	"island/serializer"
)

type CollectionUpdateService struct {
	ArticleID  uint `json:"articleID" binding:"required"` // 文章ID
	Collection bool `json:"collection"`                   // 收藏
}

func (service *CollectionUpdateService) Update(userId uint) *serializer.ErrorResponse {

	if service.Collection == true {
		return service.Create(userId)
	} else {
		return service.Delete(userId)
	}
}

func (service *CollectionUpdateService) Create(userId uint) *serializer.ErrorResponse {
	collection := model.Collection{
		ArticleID: service.ArticleID,
		CreatorID: userId,
	}
	if err := model.DB.Create(&collection).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    5001,
			Message: "收藏失败",
			Error:   err.Error(),
		}
	}
	return nil
}

func (service *CollectionUpdateService) Delete(userId uint) *serializer.ErrorResponse {
	collection := model.Collection{
		ArticleID: service.ArticleID,
		CreatorID: userId,
	}
	if err := model.DB.First(&collection).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    4004,
			Message: "该用户未收藏该文章",
			Error:   err.Error(),
		}
	}
	if err := model.DB.Delete(&collection).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    5001,
			Message: "取消收藏失败",
			Error:   err.Error(),
		}
	}
	return nil
}
