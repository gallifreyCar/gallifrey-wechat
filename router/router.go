package router

import (
	"github.com/gallifreyCar/gallifrey-wechat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.Ping)
	r.POST("/user", service.CreateUser)
	return r
}
