package v1

var UserCtrl *UserController

func Init() {
	UserCtrl = NewUserController()

}
