package apiv1

import (
	"board4486/httpserver/protocol/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProtocolHandler(controllerHandler func(ctx *common.ApiContext)) func(c *gin.Context) {
	return func(c *gin.Context) {
		ac := common.ApiContext{
			Request:  common.BuildRequest(c),
			Response: common.BuildEmptyResponse(c),
		}
		ac.Response.Version = "v1"
		controllerHandler(&ac)

		c.JSON(http.StatusOK, ac.Response)
	}
}
