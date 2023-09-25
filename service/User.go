package service

import (
	"fmt"
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

func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.PostForm("username")
	user.Password = c.PostForm("password")
	database.UpdateUser(&user)
	c.JSONP(200, gin.H{"message": "success"})
}

func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.PostForm("username")
	database.DeleteUser(&user)
	c.JSONP(200, gin.H{"message": "success"})
}

func GetUser(c *gin.Context) {
	identity := c.Param("id")
	user := models.UserBasic{}
	user.Identity = identity
	database.GetUser(&user)
	c.JSONP(200, gin.H{"message": "success", "data": user})
}

func GetUsers(c *gin.Context) {
	var users []models.UserBasic
	database.GetUsers(&users)
	fmt.Println(users)
	c.JSONP(200, gin.H{"message": "success", "data": users})
}
