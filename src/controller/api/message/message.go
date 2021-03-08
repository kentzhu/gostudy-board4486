package message

import (
	"board4486/controller/api"
	"board4486/httpserver/protocol/common"
	"board4486/httpserver/protocol/common/code"
	"board4486/service/message"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
)

func List(c *common.ApiContext) {
	c.Response.SuccessWithData(gin.H{"list": message.List()})
}

func Create(c *common.ApiContext) {
	var newMessage message.Message
	err := c.Request.GinContext.ShouldBindBodyWith(&newMessage, binding.JSON)
	if err != nil {
		c.Response.FailWithError(err)
		return
	}
	err = message.Insert(&newMessage)
	c.Response.SuccessWithData(gin.H{})
}

func Remove(c *common.ApiContext) {
	// 检查登录
	if !api.IsAuthValid(c) {
		c.Response.FailWithCode(code.AuthRequired)
		return
	}

	type payload struct {
		MessageId int `json:"id"`
	}
	var input payload
	err := c.Request.GinContext.ShouldBindBodyWith(&input, binding.JSON)
	if err != nil {
		c.Response.FailWithError(err)
		return
	}
	err = message.Delete(&input.MessageId)
	if err != nil {
		c.Response.FailWithError(err)
		return
	}
	c.Response.SuccessWithData(gin.H{})
}
