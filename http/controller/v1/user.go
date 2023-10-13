package v1

import (
	"fmt"
	"github.com/gallifreyCar/gallifrey-wechat/database"
	"github.com/gallifreyCar/gallifrey-wechat/models"
	"github.com/gallifreyCar/gallifrey-wechat/server"
	"github.com/gallifreyCar/gallifrey-wechat/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type UserController struct {
	Service *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		Service: &service.UserService{
			Db: &database.UserDao{},
		},
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	user := models.UserBasic{
		Name:     c.PostForm("username"),
		Password: c.PostForm("password"),
	}
	if err := u.Service.CreateUser(&user); err != nil {
		c.JSONP(200, gin.H{"message": err.Error()})
		return
	}
	c.JSONP(200, gin.H{"message": "success"})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	err := u.Service.UpdateUser(c.PostForm("id"), c.PostForm("username"), c.PostForm("password"))
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSONP(200, gin.H{"message": "success"})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	err := u.Service.DeleteUser(c.Param("id"))
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSONP(200, gin.H{"message": "success"})
}

func (u *UserController) GetUserByID(c *gin.Context) {
	user, err := u.Service.GetUserByID(c.Param("id"))
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSONP(200, gin.H{"message": "success", "data": user})
}

func (u *UserController) GetUsers(c *gin.Context) {
	users, err := u.Service.GetUsers()
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSONP(200, gin.H{"message": "success", "data": users})
}

var wsUP = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域
	},
}

func (u *UserController) Chat(c *gin.Context) {
	// 升级webSocket协议
	ws, err := wsUP.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	name := c.Query("name")
	client := &server.Client{
		Name:      name,
		Conn:      ws,
		DataQueue: make(chan []byte),
	}

	server.MyServer.Clients[name] = client

	// 读取数据
	go client.Read()
	// 发送数据
	go client.Write()

}
