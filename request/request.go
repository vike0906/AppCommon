package request

import "time"

type CommonResponse struct {
	Code      int
	Message   string
	Body      interface{}
	timestamp int64
}

func Success() CommonResponse {
	return CommonResponse{0, "success", "", time.Now().Unix()}
}

func Error(errorInfo string) CommonResponse {
	return CommonResponse{0, errorInfo, "", time.Now().Unix()}
}

func Response(body interface{}) CommonResponse {
	return CommonResponse{0, "success", body, time.Now().Unix()}
}
