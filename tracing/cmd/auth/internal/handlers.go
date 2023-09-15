package internal

import (
	"github.com/gin-gonic/gin"
	"tracing/pkg/common/api"
)

func GenTokenHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req TokenGenReq
		err := context.ShouldBindJSON(&req)
		if err != nil {
			context.JSON(200, TokenGenResp{
				CommonResp: api.BuildCommonRespFailed(),
			})
			return
		}

		token := GenToken(req.AccountID)
		context.JSON(200, TokenGenResp{
			CommonResp: api.BuildCommonRespSuccess(),
			Token:      token,
		})
	}
}

func DeleteTokenHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req DeleteTokenReq
		err := context.ShouldBindJSON(&req)
		if err != nil {
			context.JSON(200, DeleteTokenResp{
				CommonResp: api.BuildCommonRespFailed(),
			})
			return
		}
		DeleteToken(req.AccountID)
		context.JSON(200, DeleteTokenResp{
			CommonResp: api.BuildCommonRespSuccess(),
		})
	}
}

func VerifyTokenHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req AuthVerifyReq
		err := context.ShouldBindJSON(&req)
		if err != nil {
			context.JSON(200, AuthVerifyResp{
				CommonResp: api.BuildCommonRespFailed(),
			})
			return
		}

		pass := TokenVerification(req.AccountID, req.Token)
		if pass {
			context.JSON(200, AuthVerifyResp{
				CommonResp: api.BuildCommonRespSuccess(),
			})
		} else {
			context.JSON(200, AuthVerifyResp{
				CommonResp: api.BuildCommonRespFailed(),
				ExtraMsg:   "invalid token",
			})
		}
	}
}
