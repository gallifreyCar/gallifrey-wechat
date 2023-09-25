package service

import (
	"github.com/gallifreyCar/gallifrey-wechat/database"
	"github.com/gallifreyCar/gallifrey-wechat/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.PostForm("username")
	user.Password = c.PostForm("password")
	database.CreateUser(&user)
	c.JSONP(200, gin.H{"message": "success"})
}
