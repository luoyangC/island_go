package middleware

import (
	"island/serializer"
	"island/utils"
	"log"

	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(200, serializer.ErrorResponse{
				Code:    4003,
				Message: "需要登录",
			})
			c.Abort()
			return
		}
		log.Print("get token: ", token)
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(200, serializer.ErrorResponse{
				Code:    5000,
				Message: "解析Token失败",
				Error:   err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
