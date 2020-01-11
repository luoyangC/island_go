package api

import (
	"github.com/gin-gonic/gin"

	"island/service"
)

// @Summary 发送短信验证码
// @Tags 通用
// @Accept  json
// @Produce json
// @Param code body service.SendCodeService true "验证码"
// @Success 200 {object} serializer.Response
// @Router /api/v1/code [post]
func SendCode(c *gin.Context) {
	var s service.SendCodeService
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	if err := s.SendCode(); err != nil {
		c.JSON(200, err)
		return
	}
	c.JSON(200, Response(nil))
}
