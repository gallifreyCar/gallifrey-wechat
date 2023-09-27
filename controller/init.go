package controller

var UserCtrl *UserController

func Init() {
	UserCtrl = NewUserController()

}
