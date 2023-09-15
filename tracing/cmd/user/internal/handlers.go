package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"tracing/pkg/common/api"
	"tracing/pkg/registry"
)

func LoginHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req LoginReq
		err := context.ShouldBindJSON(&req)
		if err != nil {
			context.JSON(200, LoginResp{
				CommonResp: api.BuildCommonRespFailed(),
				ExtraMsg:   fmt.Sprintf("err: %v", err),
			})
			return
		}

		match := req.AccountID == "leebai" && req.Password == "123456"
		if !match {
			context.JSON(200, LoginResp{
				CommonResp: api.BuildCommonRespFailed(),
				ExtraMsg:   "accountID or password error",
			})
		}

		resp, err := resty.New().R().
			SetBody(map[string]interface{}{"accountID": "leebai", "password": "123456"}).
			Post(fmt.Sprintf("%v%v", registry.AuthServiceURL(), "/genToken"))
		if err != nil {
			context.JSON(200, LoginResp{
				CommonResp: api.BuildCommonRespFailed(),
				ExtraMsg:   fmt.Sprintf("%v", err),
			})
			return
		}
		fmt.Printf("LoginHandler /genToken resp: %v\n", resp)

		context.JSON(200, LoginResp{
			CommonResp: api.BuildCommonRespSuccess(),
		})
	}
}

func LogoutHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req LogoutReq
		err := context.ShouldBindJSON(&req)
		if err != nil {
			context.JSON(200, LogoutResp{
				CommonResp: api.BuildCommonRespFailed(),
				ExtraMsg:   fmt.Sprintf("err: %v", err),
			})
			return
		}

		resp, err := resty.New().R().
			SetBody(map[string]interface{}{"accountID": req.AccountID}).
			Post(fmt.Sprintf("%v%v", registry.AuthServiceURL(), "/deleteToken"))
		if err != nil {
			context.JSON(200, LogoutResp{
				CommonResp: api.BuildCommonRespFailed(),
				ExtraMsg:   fmt.Sprintf("%v", err),
			})
			return
		}
		fmt.Printf("LogoutHandler /deleteToken resp: %v\n", resp)

		context.JSON(200, LogoutResp{
			CommonResp: api.BuildCommonRespSuccess(),
		})
	}
}
