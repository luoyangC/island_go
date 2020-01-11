package serializer

import (
	"island/model"
)

// 用户序列化器
type User struct {
	ID        uint   `form:"id" json:"id"`         // 用户ID
	UserName  string `form:"name" json:"username"` // 用户名称
	Status    int    `json:"status"`               // 用户状态
	Avatar    string `json:"avatar"`               // 用户头像
	Email     string `json:"email"`                // 用户邮箱
	CreatedAt int64  `json:"createdAt"`            // 创建时间
}

// 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"` // 用户信息
}

// 用户列表序列化
type UserListResponse struct {
	Response
	Count int `json:"count"`
	Data []User `json:"data"`
}

// 序列化用户响应
func BuildUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		Response: Response{
			Code:    2000,
			Message: "success",
		},
		Data: User{
			ID:        user.ID,
			UserName:  user.UserName,
			Status:    user.Status,
			Avatar:    user.Avatar,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Unix(),
		},
	}
}

func BuildUserListResponse(items []model.User, count int) *UserListResponse {
	var users []User
	for _, item := range items {
		user := User{
			ID:        item.ID,
			UserName:  item.UserName,
			Status:    item.Status,
			Avatar:    item.Avatar,
			Email:     item.Email,
			CreatedAt: item.CreatedAt.Unix(),
		}
		users = append(users, user)
	}
	return &UserListResponse{
		Response: Response{
			Code:    2000,
			Message: "success",
		},
		Count:count,
		Data: users,
	}
}
