package serializer

import "island/model"

type Message struct {
	Content  string `json:"content"`  // 留言内容
	Avatar   string `json:"avatar"`   // 用户头像
	NickName string `json:"nickName"` // 用户昵称
}

type MessageResponse struct {
	Response
	Data *Message `json:"data"`
}

type MessageListResponse struct {
	Response
	Count int        `json:"count"`
	Data  []*Message `json:"data"`
}

func BuildMessageResponse(message *model.Message) *MessageResponse {
	return &MessageResponse{
		Response: Response{
			Code:    2000,
			Message: "success",
		},
		Data: &Message{
			Content:  message.Content,
			Avatar:   message.Avatar,
			NickName: message.NickName,
		},
	}
}

func BuildMessageListResponse(messages []*Message, count int) *MessageListResponse {
	return &MessageListResponse{
		Response: Response{
			Code:    2000,
			Message: "success",
		},
		Count: count,
		Data:  messages,
	}
}
