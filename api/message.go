package api

import (
	"github.com/gin-gonic/gin"
	"island/serializer"
	"island/service"
)

// @Summary 添加留言
// @Tags 留言
// @Accept  json
// @Produce json
// @Param article body service.MessageCreateService true "文章"
// @Success 200 {object} serializer.MessageResponse
// @Router /api/v1/message [post]
func MessageCreate(c *gin.Context) {
	var s service.MessageCreateService
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	if message, err := s.Create(c); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildMessageResponse(message)
		c.JSON(200, res)
	}
}

// @Summary 留言列表
// @Tags 留言
// @Accept  json
// @Produce json
// @Param limit query int false "每页数量"
// @Param page query int false "当前页数"
// @Success 200 {object} serializer.MessageListResponse
// @Router /api/v1/messages [get]
func MessageList(c *gin.Context)  {
	var s service.MessageListService
	if messages, count, err := s.List(c); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildMessageListResponse(messages, count)
		c.JSON(200, res)
	}
}