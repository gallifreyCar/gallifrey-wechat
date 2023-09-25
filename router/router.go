package router

import (
	"github.com/gallifreyCar/gallifrey-wechat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	ug := r.Group("/user")
	{
		ug.POST("/create", service.CreateUser)
		ug.PUT("/update", service.UpdateUser)
		ug.DELETE("/delete", service.DeleteUser)
		ug.GET("/get/:id", service.GetUser)
		ug.GET("/gets", service.GetUsers)
	}
	return r
}
