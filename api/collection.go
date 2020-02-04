package api

import (
	"github.com/gin-gonic/gin"

	"island/serializer"
	"island/service"
	"island/utils"
)

// @Summary 设置或取消收藏
// @Tags 收藏
// @Accept  json
// @Produce json
// @Param article body service.CollectionUpdateService true "文章"
// @Success 200 {object} serializer.MessageResponse
// @Router /api/v1/collection [put]
// @Security ApiKeyAuth
func CollectionUpdate(c *gin.Context) {
	var s service.CollectionUpdateService
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

// @Summary 收藏列表
// @Tags 收藏
// @Accept  json
// @Produce json
// @Param id path int true "ID"
// @Param limit query int false "每页数量"
// @Param page query int false "当前页数"
// @Success 200 {object} serializer.CollectionListResponse
// @Router /api/v1/collection/{id} [get]
func CollectionList(c *gin.Context)  {
	var s service.CollectionListService

	limit, offset, err := utils.Pagination(c)
	if err != nil {
		c.JSON(200, err)
		return
	}

	userId, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(200, err)
		return
	}

	if collections, count, err := s.List(limit, offset, userId); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildCollectionListResponse(collections, count)
		c.JSON(200, res)
	}
}