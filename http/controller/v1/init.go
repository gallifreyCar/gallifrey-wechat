package v1

var UserCtrl *UserController
var AuthCtrl *AuthController

func Init() {
	UserCtrl = NewUserController()
	AuthCtrl = NewAuthController()
}
