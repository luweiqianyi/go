package internal

import (
	"tracing/pkg/common/api"
)

type (
	LoginReq struct {
		AccountID string `json:"accountID"`
		Password  string `json:"password"`
	}

	LoginResp struct {
		api.CommonResp
		Token    string `json:"token,omitempty"`
		ExtraMsg string `json:"extraMsg,omitempty"`
	}
)

type (
	LogoutReq struct {
		AccountID string `json:"accountID"`
	}

	LogoutResp struct {
		api.CommonResp
		ExtraMsg string `json:"extraMsg,omitempty"`
	}
)
