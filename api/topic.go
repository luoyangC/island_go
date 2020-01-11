package api

import (
	"github.com/gin-gonic/gin"

	"island/serializer"
	"island/service"
	"island/utils"
)

// @Summary 新建话题
// @Tags 话题
// @Accept  json
// @Produce json
// @Param topic body service.TopicCreateService true "话题"
// @Success 200 {object} serializer.TopicResponse
// @Router /api/v1/topic [post]
// @Security ApiKeyAuth
func TopicCreate(c *gin.Context) {
	var s service.TopicCreateService
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	claims := c.MustGet("claims").(*utils.CustomClaims)
	if topic, err := s.Create(claims.ID); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildTopicResponse(topic)
		c.JSON(200, res)
	}
}

// @Summary 话题列表
// @Tags 话题
// @Accept  json
// @Produce json
// @Param limit query int false "每页数量"
// @Param page query int false "当前页数"
// @Success 200 {object} serializer.TopicListResponse
// @Router /api/v1/topics [get]
func TopicList(c *gin.Context)  {
	var s service.TopicListService
	if topics, count, err := s.List(c); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildTopicListResponse(topics, count)
		c.JSON(200, res)
	}
}

// @Summary 话题详情
// @Tags 话题
// @Accept  json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} serializer.TopicResponse
// @Router /api/v1/topic/{id} [get]
func TopicDetail(c *gin.Context)  {
	var s service.TopicDetailService
	if topic, err := s.Get(c.Param("id")); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildTopicResponse(topic)
		c.JSON(200, res)
	}
}

// @Summary 更新话题
// @Tags 话题
// @Accept  json
// @Produce json
// @Param id path int true "ID"
// @Param update body service.TopicUpdateService true "更新"
// @Success 200 {object} serializer.TopicResponse
// @Router /api/v1/topic/{id} [put]
// @Security ApiKeyAuth
func TopicUpdate(c *gin.Context)  {
	var s service.TopicUpdateService
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	claims := c.MustGet("claims").(*utils.CustomClaims)
	if topic, err := s.Update(c.Param("id"), claims.ID); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildTopicResponse(topic)
		c.JSON(200, res)
	}
}

// @Summary 删除话题
// @Tags 话题
// @Accept  json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} serializer.Response
// @Router /api/v1/topic/{id} [delete]
// @Security ApiKeyAuth
func TopicDelete(c *gin.Context)  {
	var s service.TopicDeleteService
	claims := c.MustGet("claims").(*utils.CustomClaims)
	if err := s.Delete(c.Param("id"), claims.ID); err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, Response(nil))
	}
}