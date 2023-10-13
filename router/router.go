package router

import (
	"github.com/gallifreyCar/gallifrey-wechat/http/controller/v1"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	ug := r.Group("/api/v1")
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

func Handle404Route(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    "404",
			"message": "路由不存在,请检查请求方法和请求路径",
		})
	})
}
