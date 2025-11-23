package controller

import (
	"task4/internal/dto"
	"task4/internal/service"
	"task4/pkg/response"

	"github.com/gin-gonic/gin"
)

// 定义请求
type UserController struct{}

// var AuthController = new(authController)
func (uc *UserController) GetUserInfo(c *gin.Context) {
	req := &dto.RegisterRequest{}
	// 参数校验
	if err := c.ShouldBindJSON(req); err != nil {
		response.BadRequest(c, "参数校验失败")
		return
	}
	// service层
	if error := service.UserService.Register(req); error != nil {
		response.InternalServerError(c, error.Error())
		return
	}

	response.Success(c, "success")
}
