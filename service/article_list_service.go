package service

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"island/model"
	"island/serializer"
	"island/utils"
)

type ArticleListService struct{}

func (service *ArticleListService) List(c *gin.Context) ([]*serializer.ArticleItem, int, *serializer.ErrorResponse) {
	var count int
	limit, offset, err := utils.Pagination(c)
	if err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    4000,
			Message: "错误的分页",
			Error:   err.Error(),
		}
	}
	if err := model.DB.Model(model.Article{}).Count(&count).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	var articles []*serializer.ArticleItem
	sql := "select a.title, a.image, a.tags, a.profile, a.created_at, t.title as topic_name, c.user_name as creator_name, c.avatar creator_avatar " +
		"from articles a, topics t, users c " +
		"where a.deleted_at is NULL and t.deleted_at is NULL and c.deleted_at is NULL and t.id=a.topic_id and c.id=a.creator_id " +
		"limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)
	if err := model.DB.Raw(sql).Scan(&articles).Error; err != nil {
		return nil, 0, &serializer.ErrorResponse{
			Code:    5001,
			Message: "数据库查询错误",
			Error:   err.Error(),
		}
	}
	return articles, count, nil
}
