package service

import (
	"github.com/gin-gonic/gin"
	"island/model"
	"island/serializer"
)

type MessageCreateService struct {
	Content  string `json:"content" binding:"required"`  // 留言内容
	Email    string `json:"email"`                       // 邮箱
	Mobile   string `json:"mobile"`                      // 电话
	NickName string `json:"nickName" binding:"required"` // 用户昵称
}

func (service *MessageCreateService) Create(c *gin.Context) (*model.Message, *serializer.ErrorResponse) {

	message := model.Message{
		Content:  service.Content,
		Email:    service.Email,
		Mobile:   service.Mobile,
		NickName: service.NickName,
		ClientIP: c.ClientIP(),
	}

	if err := model.DB.Create(&message).Error; err != nil {
		return nil, &serializer.ErrorResponse{
			Code:    5001,
			Message: "添加留言失败",
			Error:   err.Error(),
		}
	}

	return &message, nil
}
