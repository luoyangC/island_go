package serializer

import "island/model"

type Topic struct {
	Title string `json:"title"` // 标题
	Info  string `json:"info"`  // 详情
	Icon  string `json:"icon"`  // 图标
	Image string `json:"image"` // 背景
}

type TopicResponse struct {
	Response
	Data *Topic `json:"data"`
}

type TopicListResponse struct {
	Response
	Count int     `json:"count"`
	Data  []*Topic `json:"data"`
}

func BuildTopicResponse(topic *model.Topic) *TopicResponse {
	return &TopicResponse{
		Response: Response{
			Code:    2000,
			Message: "success",
		},
		Data: &Topic{
			Title: topic.Title,
			Icon:  topic.Icon,
			Info:  topic.Info,
			Image: topic.Image,
		},
	}
}

func BuildTopicListResponse(items []*model.Topic, count int) *TopicListResponse {
	var topics []*Topic
	for _, item := range items {
		topic := &Topic{
			Title: item.Title,
			Info:  item.Info,
			Icon:  item.Icon,
			Image: item.Image,
		}
		topics = append(topics, topic)
	}
	return &TopicListResponse{
		Response: Response{
			Code:    2000,
			Message: "success",
		},
		Count: count,
		Data:  topics,
	}
}
