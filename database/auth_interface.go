package database

type AuthDB interface {
	// IsEmailExist  邮箱是否存在
	IsEmailExist(email string) bool
	// IsPhoneExist IsPhoneExist 手机号是否存在
	IsPhoneExist(phone string) bool
}
