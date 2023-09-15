package internal

import (
	"tracing/pkg/common/api"
)

type (
	AuthVerifyReq struct {
		AccountID string `json:"accountID"`
		Token     string `json:"token"`
	}

	AuthVerifyResp struct {
		api.CommonResp
		ExtraMsg string `json:"extraMsg,omitempty"`
	}
)

type (
	TokenGenReq struct {
		AccountID string `json:"accountID"`
	}

	TokenGenResp struct {
		api.CommonResp
		Token string `json:"token,omitempty"`
	}
)

type (
	DeleteTokenReq struct {
		AccountID string `json:"accountID"`
	}

	DeleteTokenResp struct {
		api.CommonResp
	}
)
