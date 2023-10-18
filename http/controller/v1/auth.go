package v1

import (
	"github.com/gallifreyCar/gallifrey-wechat/database"
	"github.com/gallifreyCar/gallifrey-wechat/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	Service *service.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		Service: &service.AuthService{
			Db: &database.AuthDao{},
		},
	}
}

// IsEmailExist 邮箱是否存在
func (a *AuthController) IsEmailExist(c *gin.Context) {
	type Email struct {
		Email string `json:"email"`
	}
	var email Email
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"exist": a.Service.IsEmailExist(email.Email)})
}

// IsPhoneExist 手机号是否存在
func (a *AuthController) IsPhoneExist(c *gin.Context) {
	type Phone struct {
		Phone string `json:"phone"`
	}
	var phone Phone
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"exist": a.Service.IsPhoneExist(phone.Phone)})
}

// TestPanics 测试panic
func (a *AuthController) TestPanics(c *gin.Context) {
	panic("test panic")
}
