package api

import (
	"github.com/gin-gonic/gin"

	"island/serializer"
	"island/service"
	"island/utils"
)

// @Summary 新建文章
// @Tags 文章
// @Accept  json
// @Produce json
// @Param article body service.ArticleCreateService true "文章"
// @Success 200 {object} serializer.ArticleResponse
// @Router /api/v1/article [post]
// @Security ApiKeyAuth
func ArticleCreate(c *gin.Context) {
	var s service.ArticleCreateService
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	claims := c.MustGet("claims").(*utils.CustomClaims)
	if article, err := s.Create(claims.ID); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildArticleResponse(article)
		c.JSON(200, res)
	}
}

// @Summary 文章列表
// @Tags 文章
// @Accept  json
// @Produce json
// @Param limit query int false "每页数量"
// @Param page query int false "当前页数"
// @Success 200 {object} serializer.ArticleListResponse
// @Router /api/v1/articles [get]
func ArticleList(c *gin.Context)  {
	var s service.ArticleListService
	if articles, count, err := s.List(c); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildArticleListResponse(articles, count)
		c.JSON(200, res)
	}
}

// @Summary 文章详情
// @Tags 文章
// @Accept  json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} serializer.ArticleResponse
// @Router /api/v1/article/{id} [get]
func ArticleDetail(c *gin.Context)  {
	var s service.ArticleDetailService
	if article, err := s.Get(c.Param("id")); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildArticleResponse(article)
		c.JSON(200, res)
	}
}

// @Summary 更新文章
// @Tags 文章
// @Accept  json
// @Produce json
// @Param id path int true "ID"
// @Param update body service.ArticleUpdateService true "更新"
// @Success 200 {object} serializer.ArticleResponse
// @Router /api/v1/article/{id} [put]
// @Security ApiKeyAuth
func ArticleUpdate(c *gin.Context)  {
	var s service.ArticleUpdateService
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	claims := c.MustGet("claims").(*utils.CustomClaims)
	if article, err := s.Update(c.Param("id"), claims.ID); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildArticleResponse(article)
		c.JSON(200, res)
	}
}

// @Summary 删除文章
// @Tags 文章
// @Accept  json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} serializer.Response
// @Router /api/v1/article/{id} [delete]
// @Security ApiKeyAuth
func ArticleDelete(c *gin.Context)  {
	var s service.TopicDeleteService
	claims := c.MustGet("claims").(*utils.CustomClaims)
	if err := s.Delete(c.Param("id"), claims.ID); err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, Response(nil))
	}
}