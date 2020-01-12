package utils

import (
	"gopkg.in/go-playground/validator.v9"
	"island/model"
)

func TopicValid(fl validator.FieldLevel) bool {

	var count int

	model.DB.Model(&model.Topic{}).Where("id=?", fl.Field().Uint()).Count(&count)
	if count <= 0 {
		return false
	}

	return true
}