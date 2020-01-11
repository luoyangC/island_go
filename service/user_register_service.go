package service

import (
	"strings"

	"island/cache"
	"island/model"
	"island/serializer"
)

// 用户注册服务
type UserRegisterService struct {
	Mobile   string `json:"mobile" binding:"required,min=11,max=11"`  // 手机号码
	Code     string `json:"code" binding:"required,min=6,max=6"`      // 验证码
	Password string `json:"password" binding:"required,min=8,max=40"` // 登录密码
}

// 用户注册验证
func (service *UserRegisterService) Valid() *serializer.ErrorResponse {

	count := 0
	model.DB.Model(&model.User{}).Where("mobile = ?", service.Mobile).Count(&count)
	if count > 0 {
		return &serializer.ErrorResponse{
			Code:    4120,
			Message: "用户手机号已注册",
		}
	}

	codeCache, err := cache.Client.Get(service.Mobile).Result()
	if err != nil {
		return &serializer.ErrorResponse{
			Code:    4121,
			Message: "无效验证码",
		}
	}

	codeList := strings.Split(codeCache, ",")
	if codeList[0] != "1" {
		return &serializer.ErrorResponse{
			Code:    4122,
			Message: "验证码类型错误",
		}
	}

	if codeList[1] != service.Code {
		return &serializer.ErrorResponse{
			Code:    4123,
			Message: "验证码错误",
		}
	}

	return nil
}

// 用户注册
func (service *UserRegisterService) Register() (*model.User, *serializer.ErrorResponse) {

	user := model.User{
		UserName: "用户" + service.Mobile,
		Mobile:   service.Mobile,
		Status:   model.Active,
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return nil, err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return &user, &serializer.ErrorResponse{
			Code:    4124,
			Message: "密码加密失败",
			Error:   err.Error(),
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return &user, &serializer.ErrorResponse{
			Code:    4125,
			Message: "注册失败",
			Error:   err.Error(),
		}
	}

	return &user, nil
}
