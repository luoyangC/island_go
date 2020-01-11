package service

import (
	"fmt"
	"island/utils"
	"strings"

	"island/cache"
	"island/model"
	"island/serializer"
)

// 用户登录的服务
type UserLoginService struct {
	Mobile   string `json:"mobile" binding:"required,min=5,max=30"` // 手机号
	Code     string `json:"code"`                                   // 验证码
	Password string `json:"password"`                               // 密码
}

// Login 用户登录函数
func (service *UserLoginService) Login() (string, *serializer.ErrorResponse) {

	if service.Mobile == "" && service.Code == "" {
		return "", &serializer.ErrorResponse{
			Code:    4110,
			Message: "参数错误",
		}
	}

	var user model.User
	if err := model.DB.Where("mobile = ?", service.Mobile).First(&user).Error; err != nil {
		return "", &serializer.ErrorResponse{
			Code:    4110,
			Message: "未找到该手机号账户",
			Error:   fmt.Sprint(err),
		}
	}

	if service.Code != "" {
		if err := checkCode(service.Mobile, service.Code); err != nil {
			return "", err
		}
	}

	if service.Password != "" {
		if user.CheckPassword(service.Password) == false {
			return "", &serializer.ErrorResponse{
				Code:    4111,
				Message: "账号或密码错误",
			}
		}
	}

	if token, err := utils.GenerateToken(&user); err != nil {
		return "", &serializer.ErrorResponse{
			Code:    4112,
			Message: "登录失败",
			Error:   fmt.Sprint(err),
		}
	} else {
		return token, nil
	}
}

// 校验验证码
func checkCode(mobile string, code string) *serializer.ErrorResponse {

	codeCache, err := cache.Client.Get(mobile).Result()
	if err != nil {
		return &serializer.ErrorResponse{
			Code:    4121,
			Message: "无效验证码",
		}
	}

	codeList := strings.Split(codeCache, ",")
	if codeList[0] != "2" {
		return &serializer.ErrorResponse{
			Code:    4122,
			Message: "验证码类型错误",
		}
	}

	if codeList[1] != code {
		return &serializer.ErrorResponse{
			Code:    4123,
			Message: "验证码错误",
		}
	}

	return nil
}
