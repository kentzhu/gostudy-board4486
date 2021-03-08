package auth

import (
	"board4486/httpserver/protocol/common"
	"board4486/service/accesstoken"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Login(ac *common.ApiContext) {
	type Input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var payload Input
	err := ac.Request.GinContext.ShouldBindBodyWith(&payload, binding.JSON)
	if err != nil {
		ac.Response.FailWithError(err)
		return
	}

	token := accesstoken.Gen(payload.Username)
	ac.Response.SuccessWithData(gin.H{
		"token": token,
	})
}

func Logout(ac *common.ApiContext) {
	ac.Response.SuccessWithData(gin.H{})
}
