package api

import (
	"encoding/json"
	"fmt"

	"gopkg.in/go-playground/validator.v8"

	"island/serializer"
)

// 正常响应
func Response(data interface{}) serializer.Response {
	return serializer.Response{
		Code:    2000,
		Data:    data,
		Message: "success",
	}
}

// 错误响应
func ErrorResponse(err error) serializer.ErrorResponse {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			return serializer.ErrorResponse{
				Code:    4000,
				Data:    nil,
				Message: fmt.Sprintf("%s%s", e.Field, e.Tag),
				Error:   err.Error(),
			}
		}
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ErrorResponse{
			Code:    4000,
			Data:    nil,
			Message: "JSON类型不匹配",
			Error:   err.Error(),
		}
	}

	return serializer.ErrorResponse{
		Code:    9999,
		Data:    nil,
		Message: "未知错误",
		Error:   err.Error(),
	}
}
