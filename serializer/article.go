package serializer

import (
	"island/model"
	"time"
)

type Article struct {
	Title     string    `json:"title" binding:"required"`   // 标题
	Image     string    `json:"image"`                      // banner图
	Tags      string    `json:"tags" binding:"required"`    // 标签
	Profile   string    `json:"profile" binding:"required"` // 简介
	Content   string    `json:"content" binding:"required"` // 内容
	CreatedAt time.Time `json:"createdAt"`                  // 创建时间
	Topic     *Topic    `json:"topic"`                      // 话题
	Creator   *User     `json:"creator"`                    // 作者
}

type ArticleItem struct {
	Title         string    `json:"title" binding:"required"`   // 标题
	Image         string    `json:"image"`                      // banner图
	Tags          string    `json:"tags" binding:"required"`    // 标签
	Profile       string    `json:"profile" binding:"required"` // 简介
	CreatedAt     time.Time `json:"createdAt"`                  // 创建时间
	TopicName     string    `json:"topicName"`                  // 话题名称
	CreatorName   string    `json:"creatorName"`                // 作者名字
	CreatorAvatar string    `json:"creatorAvatar"`              // 作者头像
}

type ArticleResponse struct {
	Response
	Data *Article `json:"data"`
}

type ArticleListResponse struct {
	Response
	Count int            `json:"count"`
	Data  []*ArticleItem `json:"data"`
}

func BuildArticleResponse(article *model.Article) *ArticleResponse {
	return &ArticleResponse{
		Response: Response{
			Code:    2000,
			Message: "success",
		},
		Data: &Article{
			Title:     article.Title,
			Image:     article.Image,
			Tags:      article.Tags,
			Profile:   article.Profile,
			Content:   article.Content,
			CreatedAt: article.CreatedAt,
			Topic: &Topic{
				Title: article.Topic.Title,
				Icon:  article.Topic.Icon,
				Info:  article.Topic.Info,
				Image: article.Topic.Image,
			},
			Creator: &User{
				ID:        article.Creator.ID,
				UserName:  article.Creator.UserName,
				Status:    article.Creator.Status,
				Avatar:    article.Creator.Avatar,
				Email:     article.Creator.Email,
				CreatedAt: article.Creator.CreatedAt,
			},
		},
	}
}

func BuildArticleListResponse(articles []*ArticleItem, count int) *ArticleListResponse {
	return &ArticleListResponse{
		Response: Response{
			Code:    2000,
			Message: "success",
		},
		Count: count,
		Data:  articles,
	}
}
