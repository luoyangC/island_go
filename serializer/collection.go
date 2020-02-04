package serializer

type Collection struct {
	ArticleID string `json:"ArticleID"` // 文章ID
	Title     string `json:"title"`     // 标题
	Profile   string `json:"profile"`   // 简介
}

type CollectionListResponse struct {
	Response
	Count int           `json:"count"`
	Data  []*Collection `json:"data"`
}

func BuildCollectionListResponse(items []*Collection, count int) *CollectionListResponse {
	return &CollectionListResponse{
		Response: Response{
			Code:    2000,
			Message: "success",
		},
		Count: count,
		Data:  items,
	}
}
