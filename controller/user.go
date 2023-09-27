package controller

import (
	"github.com/gallifreyCar/gallifrey-wechat/models"
	"github.com/gallifreyCar/gallifrey-wechat/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController() *UserController {
	svc := service.NewUserService()
	return &UserController{service: svc}
}

func (u *UserController) CreateUser(c *gin.Context) {
	user := models.UserBasic{
		Name:     c.PostForm("username"),
		Password: c.PostForm("password"),
	}
	if err := u.service.CreateUser(&user); err != nil {
		c.JSONP(200, gin.H{"message": err.Error()})
		return
	}
	c.JSONP(200, gin.H{"message": "success"})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	err := u.service.UpdateUser(c.PostForm("id"), c.PostForm("username"), c.PostForm("password"))
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSONP(200, gin.H{"message": "success"})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	err := u.service.DeleteUser(c.Param("id"))
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSONP(200, gin.H{"message": "success"})
}

func (u *UserController) GetUserByID(c *gin.Context) {
	user, err := u.service.GetUserByID(c.Param("id"))
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSONP(200, gin.H{"message": "success", "data": user})
}

func (u *UserController) GetUsers(c *gin.Context) {
	users, err := u.service.GetUsers()
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSONP(200, gin.H{"message": "success", "data": users})
}