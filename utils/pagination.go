package utils

import (
	"island/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(ctx *gin.Context) (limitInt int, offsetInt int, res *serializer.ErrorResponse) {
	limit := ctx.DefaultQuery("limit", "10")
	page := ctx.DefaultQuery("page", "1")
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 0 {
		return 0, 0, &serializer.ErrorResponse{
			Code:    4001,
			Message: "limit参数错误",
			Error:   err.Error(),
		}
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 0 {
		return 0, 0, &serializer.ErrorResponse{
			Code:    4001,
			Message: "page参数错误",
			Error:   err.Error(),
		}
	}
	if pageInt != 0 {
		pageInt--
	}
	offsetInt = limitInt * pageInt
	return limitInt, offsetInt, nil
}
