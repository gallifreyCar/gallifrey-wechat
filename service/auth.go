package service

import "github.com/gallifreyCar/gallifrey-wechat/database"

type AuthService struct {
	Db database.AuthDB
}

func (a *AuthService) IsEmailExist(email string) bool {
	return a.Db.IsEmailExist(email)
}

func (a *AuthService) IsPhoneExist(phone string) bool {
	return a.Db.IsPhoneExist(phone)
}
