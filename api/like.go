package api

import (
	"github.com/gin-gonic/gin"

	"island/serializer"
	"island/service"
	"island/utils"
)

// @Summary 设置或取消点赞
// @Tags 点赞
// @Accept  json
// @Produce json
// @Param article body service.LikeUpdateService true "点赞"
// @Success 200 {object} serializer.MessageResponse
// @Router /api/v1/like [put]
// @Security ApiKeyAuth
func LikeUpdate(c *gin.Context) {
	var s service.LikeUpdateService
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	claims := c.MustGet("claims").(*utils.CustomClaims)
	if err := s.Update(claims.ID); err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, Response(nil))
	}
}

// @Summary 点赞列表
// @Tags 点赞
// @Accept  json
// @Produce json
// @Param id path int true "ID"
// @Param limit query int false "每页数量"
// @Param page query int false "当前页数"
// @Success 200 {object} serializer.LikeListResponse
// @Router /api/v1/like [get]
// @Security ApiKeyAuth
func LikeList(c *gin.Context)  {
	var s service.LikeListService

	limit, offset, err := utils.Pagination(c)
	if err != nil {
		c.JSON(200, err)
		return
	}

	claims := c.MustGet("claims").(*utils.CustomClaims)
	if likes, count, err := s.List(limit, offset, claims.ID); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildLikeListResponse(likes, count)
		c.JSON(200, res)
	}
}