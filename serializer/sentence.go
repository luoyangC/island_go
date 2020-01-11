package serializer

import "island/model"

type Sentence struct {
	Lines string `json:"lines"` // 句子
}

type SentenceResponse struct {
	Response
	Data *Sentence `json:"data"`
}

func BuildSentenceResponse(sentence *model.Sentence) *SentenceResponse {
	return &SentenceResponse{
		Response: Response{
			Code:    2000,
			Message: "success",
		},
		Data: &Sentence{
			Lines: sentence.Lines,
		},
	}
}
