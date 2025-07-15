package utils

// Response 统一响应结构
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
}

// NewSuccessResponse 创建成功响应
func NewSuccessResponse(data interface{}, message string) Response {
	return Response{
		Success: true,
		Data:    data,
		Message: message,
	}
}

// NewErrorResponse 创建错误响应
func NewErrorResponse(err error, message string) Response {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	return Response{
		Success: false,
		Data:    nil,
		Message: message,
		Error:   errMsg,
	}
}
