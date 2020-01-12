package service

import (
	"island/model"
	"island/serializer"
)

type ArticleDetailService struct{}

func (service *ArticleDetailService) Get(id string) (*model.Article, *serializer.ErrorResponse) {
	var article model.Article
	if err := model.DB.First(&article, id).Related(&article.Topic, "TopicId").Related(&article.Creator, "CreatorId").Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    4004,
			Message: "没有找到该文章",
			Error:   err.Error(),
		}
	}
	return &article, nil
}
