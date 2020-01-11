package service

import (
	"island/model"
	"island/serializer"
	"math/rand"
)

type RandomSentenceService struct{}

func (service *RandomSentenceService) Get() (*model.Sentence, *serializer.ErrorResponse) {
	var sentence model.Sentence
	var count int
	if err := model.DB.Model(model.Sentence{}).Count(&count).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    5000,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	if err := model.DB.First(&sentence, rand.Intn(count) + 1).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    4004,
			Message: "获取失败",
			Error:   err.Error(),
		}
	}
	return &sentence, nil
}
