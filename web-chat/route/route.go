package route

import (
	"github.com/gin-gonic/gin"
	"web-chat/controller"
	"web-chat/middleware"
	"web-chat/server"
)

func Router() *gin.Engine {
	route := gin.Default()
	route.Use(middleware.Cors())
	v1 := route.Group("/api")
	{
		v1.POST("/register", controller.RegisterHandler)
		v1.POST("/login", controller.AuthHandler)
		v1.GET("/info", middleware.JWTAuthMiddleware(),controller.HomeHandler)
	}

	// websocket
	{
		route.GET("/ws/socket", server.Websocket.Handle())
	}
	return route
}
