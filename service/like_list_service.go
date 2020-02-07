package service

import (
	"island/model"
	"island/serializer"
)

type LikeListService struct {}

func (service *LikeListService) List(limit int, offset int, userId uint) ([]*serializer.Like, int, *serializer.ErrorResponse) {
	var count int
	if err := model.DB.Model(model.Like{}).Where("creator_id = ?", userId).Count(&count).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	var likes []*serializer.Like
	if err := model.DB.Model(model.Like{}).Where("creator_id = ?", userId).Where("likes.deleted_at IS NULL").
		Limit(limit).Offset(offset).Find(&likes).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库错误",
			Error:   err.Error(),
		}
	}
	return likes, count, nil
}
