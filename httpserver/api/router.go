package api

import (
	"websocket/httpserver/api/echo"
	"websocket/httpserver/api/home"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	go func() {
		router.GET("/home", home.Home)
		router.GET("/ws", echo.Ws)
	}()
}
