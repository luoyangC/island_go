package service

import (
	"island/model"
	"island/serializer"
)

type ArticleListService struct{}

func (service *ArticleListService) List(limit int, offset int) ([]*serializer.ArticleItem, int, *serializer.ErrorResponse) {
	var count int
	if err := model.DB.Model(model.Article{}).Count(&count).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	var articles []*serializer.ArticleItem
	if err := model.DB.Table("articles").
		Select("articles.title, articles.image, articles.profile, articles.created_at, topics.title as topic_name, users.user_name AS creator_name").
		Joins("LEFT JOIN topics ON topics.id = articles.topic_id").Joins("LEFT JOIN users ON users.id = articles.creator_id").
		Limit(limit).Offset(offset).Find(&articles).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库错误",
			Error:   err.Error(),
		}
	}
	return articles, count, nil
}
