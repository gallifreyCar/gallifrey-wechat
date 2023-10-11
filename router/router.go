package router

import (
	"github.com/gallifreyCar/gallifrey-wechat/http/controller/v1"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	ug := r.Group("/user")
	{
		ug.POST("/create", v1.UserCtrl.CreateUser)
		ug.PUT("/update", v1.UserCtrl.UpdateUser)
		ug.DELETE("/delete/:id", v1.UserCtrl.DeleteUser)
		ug.GET("/get/:id", v1.UserCtrl.GetUserByID)
		ug.GET("/gets", v1.UserCtrl.GetUsers)
		ug.GET("/ws", v1.UserCtrl.Chat)
	}
	return r
}
