package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(ctx *gin.Context) (limitInt int, offsetInt int, err error) {
	limit := ctx.DefaultQuery("limit", "10")
	page := ctx.DefaultQuery("page", "1")
	limitInt, err = strconv.Atoi(limit)
	if err != nil || limitInt < 0 {
		return 0, 0, err
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 0 {
		return 0, 0, err
	}
	if pageInt != 0 {
		pageInt--
	}
	offsetInt = limitInt * pageInt
	return limitInt, offsetInt, nil
}
