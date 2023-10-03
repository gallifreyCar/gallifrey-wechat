package router

import (
	"github.com/gallifreyCar/gallifrey-wechat/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	ug := r.Group("/user")
	{
		ug.POST("/create", controller.UserCtrl.CreateUser)
		ug.PUT("/update", controller.UserCtrl.UpdateUser)
		ug.DELETE("/delete/:id", controller.UserCtrl.DeleteUser)
		ug.GET("/get/:id", controller.UserCtrl.GetUserByID)
		ug.GET("/gets", controller.UserCtrl.GetUsers)
		ug.GET("/ws", controller.UserCtrl.Chat)
	}
	return r
}
