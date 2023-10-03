package service

import (
	"fmt"
	"github.com/gallifreyCar/gallifrey-wechat/database"
	"github.com/gallifreyCar/gallifrey-wechat/models"
	"github.com/gallifreyCar/gallifrey-wechat/utils"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

type UserService struct {
	db database.UserDB
}

func NewUserService() *UserService {
	return &UserService{db: database.NewUserDao()}
}

func (s *UserService) CreateUser(user *models.UserBasic) error {
	// 检查用户名和密码是否为空
	if user.Name == "" || user.Password == "" {
		return fmt.Errorf("username or password is empty")
	}
	// 检查用户是否存在
	if _, err := s.GetUserByUsername(user.Name); err == nil {
		return fmt.Errorf("username already exists")
	}
	// 密码加密+盐
	salt := utils.RandomSalt()
	user.Salt = salt
	user.Password = utils.Md5Encode(user.Password + salt)

	// 创建用户
	if s.db.CreateUser(user).RowsAffected == 0 {
		return fmt.Errorf("create failed")
	}
	return nil
}

func (s *UserService) GetUserByUsername(username string) (*models.UserBasic, error) {
	user := models.UserBasic{Name: username}
	if s.db.GetUser(&user).RowsAffected == 0 {
		return nil, fmt.Errorf("user not exists")
	}

	return &user, nil
}

func (s *UserService) GetUserByID(id string) (*models.UserBasic, error) {
	user := models.UserBasic{Identity: id}
	if s.db.GetUser(&user).RowsAffected == 0 {
		return nil, fmt.Errorf("user not exists")
	}
	return &user, nil
}

func (s *UserService) UpdateUser(id, username, password string) error {
	// 查询用户
	user := models.UserBasic{Identity: id}
	if s.db.GetUser(&user).RowsAffected == 0 {
		return fmt.Errorf("user not exists")
	}

	// 更新用户信息
	user.Name = username
	user.Password = password

	//校验
	err := validate.Var(user.Name, "required")
	err = validate.Var(user.Password, "required")
	if err != nil {
		return err
	}

	if s.db.UpdateUser(&user).RowsAffected == 0 {
		return fmt.Errorf("update failed")
	}

	return nil
}

func (s *UserService) DeleteUser(id string) error {
	user := models.UserBasic{Identity: id}
	if s.db.DeleteUser(&user).RowsAffected == 0 {
		return fmt.Errorf("delete failed")
	}
	return nil
}

func (s *UserService) GetUsers() (*[]models.UserBasic, error) {
	var users []models.UserBasic
	s.db.GetUsers(&users)
	return &users, nil
}
