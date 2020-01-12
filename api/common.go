package api

import (
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"

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
	if _, ok := err.(validator.ValidationErrors); ok {
		return serializer.ErrorResponse{
			Code:    4000,
			Data:    nil,
			Message: "参数验证错误",
			Error:   err.Error(),
		}
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ErrorResponse{
			Code:    4000,
			Data:    nil,
			Message: "参数类型错误",
			Error:   err.Error(),
		}
	}

	if _, ok := err.(*json.SyntaxError); ok {
		return serializer.ErrorResponse{
			Code:    4000,
			Data:    nil,
			Message: "JSON格式错误",
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
