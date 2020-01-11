package service

import (
	"github.com/gin-gonic/gin"
	"island/model"
	"island/serializer"
	"island/utils"
)

type TopicListService struct {}

func (service *TopicListService) List(c *gin.Context) ([]*model.Topic, int, *serializer.ErrorResponse) {
	var topics []*model.Topic
	var count int
	limit, offset, err := utils.Pagination(c)
	if err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5000,
			Message: "错误的分页",
			Error:   err.Error(),
		}
	}
	if err := model.DB.Model(model.Topic{}).Count(&count).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5000,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	if err := model.DB.Offset(offset).Limit(limit).Find(&topics).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5000,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	return topics, count, nil
}
