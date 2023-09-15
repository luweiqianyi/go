package api

const (
	HttpRequestCodeSuccess = 0
	HttpRequestCodeFailed  = 1

	HttpRequestMsgSuccess = "success"
	HttpRequestMsgFailed  = "failed"
)

type CommonResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BuildCommonRespSuccess() CommonResp {
	return CommonResp{
		Code:    HttpRequestCodeSuccess,
		Message: HttpRequestMsgSuccess,
	}
}

func BuildCommonRespFailed() CommonResp {
	return CommonResp{
		Code:    HttpRequestCodeFailed,
		Message: HttpRequestMsgFailed,
	}
}
