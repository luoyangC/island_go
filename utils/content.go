package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"island/serializer"
)

func GetUserID(ctx *gin.Context) (userId uint, res *serializer.ErrorResponse) {

	if claims, _ := ctx.Get("claims"); claims != nil {
		userId = claims.(*CustomClaims).ID
	}
	if paramId, err := strconv.ParseUint(ctx.Param("id"), 10, 64); err == nil {
		userId = uint(paramId)
	}
	if userId == 0 {
		return 0, &serializer.ErrorResponse{
			Code:    4001,
			Message: "获取用户ID失败",
		}
	}
	return userId, nil
}