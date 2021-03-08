package common

import "github.com/gin-gonic/gin"

type ApiRequest struct {
	GinContext *gin.Context
}

func BuildRequest(ctx *gin.Context) *ApiRequest {
	req := ApiRequest{
		GinContext: ctx,
	}
	return &req
}
