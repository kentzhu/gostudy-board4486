package router

import (
	"board4486/controller/api/auth"
	"board4486/controller/api/message"
	"board4486/httpserver/protocol/apiv1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Binder(engine *gin.Engine) {

	// 绑定首页
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/static/index.html")
	})

	// 绑定静态资源
	engine.StaticFS("/static", http.Dir("static"))

	// 绑定API
	engine.POST("/api/auth/login", apiv1.ProtocolHandler(auth.Login))
	engine.POST("/api/auth/logout", apiv1.ProtocolHandler(auth.Logout))
	engine.POST("/api/message/list", apiv1.ProtocolHandler(message.List))
	engine.POST("/api/message/create", apiv1.ProtocolHandler(message.Create))
	engine.POST("/api/message/remove", apiv1.ProtocolHandler(message.Remove))

}
