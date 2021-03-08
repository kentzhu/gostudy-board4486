package auth

import (
	"board4486/httpserver/protocol/common"
	"board4486/httpserver/protocol/common/code"
	"board4486/service/accesstoken"
	"board4486/service/admin"
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

	if !admin.IsAdminLoginCorrect(payload.Username, payload.Password) {
		ac.Response.FailWithMessage(code.AuthRequired, "Login fail. Username or password was wrong.")
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
