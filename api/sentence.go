package api

import (
	"github.com/gin-gonic/gin"
	"island/serializer"
	"island/service"
)

// @Summary 随机获取定场诗
// @Tags 通用
// @Accept  json
// @Produce json
// @Success 200 {object} serializer.SentenceResponse
// @Router /api/v1/sentence [get]
func RandomSentence(c *gin.Context)  {
	var s service.RandomSentenceService
	if sentence, err := s.Get(); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildSentenceResponse(sentence)
		c.JSON(200, res)
	}
}