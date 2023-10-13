package router

import (
	v1 "github.com/gallifreyCar/gallifrey-wechat/http/controller/v1"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {

	v1ApiGroup := r.Group("/api/v1")
	{
		authGroup := v1ApiGroup.Group("/auth")
		{
			authGroup.POST("/signup/email/exist", v1.AuthCtrl.IsEmailExist)
			authGroup.POST("/signup/phone/exist", v1.AuthCtrl.IsPhoneExist)
		}

		userGroup := v1ApiGroup.Group("/user")
		{

			userGroup.POST("/create", v1.UserCtrl.CreateUser)
			userGroup.PUT("/update", v1.UserCtrl.UpdateUser)
			userGroup.DELETE("/delete/:id", v1.UserCtrl.DeleteUser)
			userGroup.GET("/get/:id", v1.UserCtrl.GetUserByID)
			userGroup.GET("/gets", v1.UserCtrl.GetUsers)
			userGroup.GET("/ws", v1.UserCtrl.Chat)
		}
	}

	return r
}

func Handle404Route(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    "404",
			"message": "路由不存在,请检查请求方法和请求路径",
		})
	})
}
