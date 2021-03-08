package common

import (
	"board4486/httpserver/protocol/common/code"
	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
	Version string                 `json:"_v"`
}

func BuildEmptyResponse(ctx *gin.Context) *ApiResponse {
	req := ApiResponse{}
	return &req
}

func (rsp *ApiResponse) Success() {
	rsp.Code = code.Success
	rsp.Data = gin.H{}
}

func (rsp *ApiResponse) SuccessWithData(data map[string]interface{}) {
	rsp.Code = code.Success
	rsp.Data = data
}

func (rsp *ApiResponse) Fail() {
	rsp.Code = code.Fail
	rsp.Message = "request fail"
	rsp.Data = gin.H{}
}

func (rsp *ApiResponse) FailWithCode(rspCode string) {
	rsp.Code = rspCode
	rsp.Message = code.GetMessage(rspCode)
	rsp.Data = gin.H{}
}

func (rsp *ApiResponse) FailWithMessage(rspCode string, message string) {
	rsp.Code = rspCode
	rsp.Message = message
	rsp.Data = gin.H{}
}

func (rsp *ApiResponse) FailWithError(err error) {
	rsp.Code = code.Fail
	rsp.Message = err.Error()
	rsp.Data = gin.H{}
}
