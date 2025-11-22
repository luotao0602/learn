package service

import (
	"task4/internal/dto"
	"task4/internal/model"
	"task4/pkg/db"
	"task4/pkg/exception"
)

// 定义结构体
type userService struct{}

// 定义全局变量，单例
var UserService = new(userService)

func (user *userService) Register(req *dto.RegisterRequest) error {
	db := db.GetGormDB()
	// 先查数据库中是否已存在
	var exsitUser model.User
	if error := db.Debug().Where("username = ?", req.Username).First(&exsitUser).Error; error == nil {
		return exception.NewSystemException("user exsits")
	}
	// 检查邮箱
	if error := db.Where("email = ?", req.Email).First(&exsitUser).Error; error == nil {
		return exception.NewSystemException("email exsits")
	}
	// 插入数据
	userDo := &model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	result := db.Debug().Create(userDo)
	if result.Error != nil {
		return exception.NewSystemException(result.Error.Error())
	}
	return nil
}

func (user *userService) Login(loginReq *dto.LoginRequest) error {
	db := db.GetGormDB()
	existUser := &model.User{}
	if err := db.Where("username = ?", loginReq.Username).First(existUser); err != nil {
		return exception.NewSystemException(err.Error.Error())
	}
	if !existUser.CheckPassword(loginReq.Password) {
		return exception.NewSystemException("password is not right")
	}
	// TODO
	return nil
}
