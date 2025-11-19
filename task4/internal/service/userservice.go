package service

import "task4/internal/model"

// 定义结构体
type userService struct{}

// 定义全局变量
var UserService = new(userService)

func (user *userService) register(userData *model.User) error {

	return nil
}
