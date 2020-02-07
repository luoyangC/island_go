package serializer

type Like struct {
	LikeType  string `json:"type"` // 类型
	LikeID    string `json:"id"`
	CreatorID string `json:"creatorId"`
}

type LikeListResponse struct {
	Response
	Count int     `json:"count"`
	Data  []*Like `json:"data"`
}

func BuildLikeListResponse(items []*Like, count int) *LikeListResponse {
	return &LikeListResponse{
		Response: Response{
			Code:    2000,
			Message: "success",
		},
		Count: count,
		Data:  items,
	}
}
