package serializer

// Response 团队基础序列化器
type Response struct {
	Code    int         `json:"code" example:"2000"`       // 状态码
	Data    interface{} `json:"data"`                      // 响应数据
	Message string      `json:"message" example:"success"` // 自定义消息
}

// TrackedErrorResponse 有追踪信息的错误响应
type ErrorResponse struct {
	Code    int         `json:"code" example:"9999"`     // 状态码
	Data    interface{} `json:"data"`                    // 响应数据
	Message string      `json:"message" example:"error"` // 自定义消息
	Error   string      `json:"error"`                   // 错误详情
}
