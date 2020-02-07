package service

import (
	"island/model"
	"island/serializer"
)

type LikeUpdateService struct {
	LikeType string `json:"type"` // 点赞类型
	LikeID   uint   `json:"id"`   // 点赞ID
	Like     bool   `json:"like"` // 点赞状态
}

func (service *LikeUpdateService) Update(userId uint) *serializer.ErrorResponse {

	if service.Like == true {
		return service.Create(userId)
	} else {
		return service.Delete(userId)
	}
}

func (service *LikeUpdateService) Create(userId uint) *serializer.ErrorResponse {
	like := model.Like{
		LikeType:  service.LikeType,
		LikeID:    service.LikeID,
		CreatorID: userId,
	}
	if err := model.DB.Unscoped().Where(&like).FirstOrCreate(&like).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    5001,
			Message: "新增点赞失败",
			Error:   err.Error(),
		}
	}
	if err := model.DB.Model(&like).Unscoped().Update("deleted_at", nil).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    5001,
			Message: "重新点赞失败",
			Error:   err.Error(),
		}
	}
	return nil
}

func (service *LikeUpdateService) Delete(userId uint) *serializer.ErrorResponse {
	like := model.Like{
		LikeType:  service.LikeType,
		LikeID:    service.LikeID,
		CreatorID: userId,
	}
	if err := model.DB.Where(&like).First(&like).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    4004,
			Message: "该用户未点赞",
			Error:   err.Error(),
		}
	}
	if err := model.DB.Delete(&like).Error; err != nil {
		return &serializer.ErrorResponse{
			Code:    5001,
			Message: "取消点赞失败",
			Error:   err.Error(),
		}
	}
	return nil
}
