package request

import "time"

type CommonResponse struct {
	Code      int
	Message   string
	Body      interface{}
	Timestamp int64
}

func Success() CommonResponse {
	return CommonResponse{Code: 0, Message: "success", Timestamp: time.Now().Unix()}
}

func Error(errorInfo string) CommonResponse {
	return CommonResponse{Code: 0, Message: errorInfo, Timestamp: time.Now().Unix()}
}

func Response(body interface{}) CommonResponse {
	return CommonResponse{Code: 0, Message: "success", Body: body, Timestamp: time.Now().Unix()}
}
