package service

import (
	"island/model"
	"island/serializer"
)

type CollectionListService struct {}

func (service *CollectionListService) List(limit int, offset int, userId uint) ([]*serializer.Collection, int, *serializer.ErrorResponse) {
	var count int
	if err := model.DB.Model(model.Collection{}).Where("creator_id = ?", userId).Count(&count).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	var collections []*serializer.Collection
	if err := model.DB.Model(model.Collection{}).Where("creator_id = ?", userId).Limit(limit).Offset(offset).Find(&collections).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库错误",
			Error:   err.Error(),
		}
	}
	return collections, count, nil
}
