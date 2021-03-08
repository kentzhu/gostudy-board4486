package api

import (
	"board4486/httpserver/protocol/common"
	"board4486/service/accesstoken"
	"github.com/gin-gonic/gin/binding"
)

func IsAuthValid(ac *common.ApiContext) bool {
	type AuthInfo struct {
		Token string `json:"_token"`
	}
	info := AuthInfo{}

	err := ac.Request.GinContext.ShouldBindBodyWith(&info, binding.JSON)
	if err != nil {
		return false
	}
	return accesstoken.IsAccessTokenValid(info.Token)
}
