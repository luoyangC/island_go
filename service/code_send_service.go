package service

import (
	"island/cache"
	"island/model"
	"island/serializer"
	"island/utils"
)

type SendCodeService struct {
	Type   int    `json:"type" binding:"required"`                 // 类型
	Mobile string `json:"mobile" binding:"required,min=11,max=11"` // 手机号码
}

func (service *SendCodeService) Valid() *serializer.ErrorResponse {

	_, err := cache.Client.Get(service.Mobile).Result()
	if err == nil {
		return &serializer.ErrorResponse{
			Code:    4100,
			Message: "距离上一次短信发送不足一分钟",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("mobile = ?", service.Mobile).Count(&count)

	if service.Type == 1 && count > 0 {
		return &serializer.ErrorResponse{
			Code:    4101,
			Message: "手机号已注册",
		}
	} else if service.Type == 2 && count == 0 {
		return &serializer.ErrorResponse{
			Code:    4102,
			Message: "该手机号未注册",
		}
	}

	return nil
}

func (service *SendCodeService) SendCode() *serializer.ErrorResponse {

	if err := service.Valid(); err != nil {
		return err
	}
	if code := utils.SendSmsCode(service.Mobile, service.Type); code != "OK" {
		return &serializer.ErrorResponse{
			Code:    4103,
			Message: code,
		}
	}
	return nil
}
