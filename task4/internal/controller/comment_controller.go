package controller

import (
	"task4/internal/dto"
	"task4/internal/service"
	"task4/pkg/response"

	"github.com/gin-gonic/gin"
)

// 定义请求
type CommentController struct{}

// var AuthController = new(authController)
func (pt *CommentController) CreateComment(c *gin.Context) {
	req := &dto.CommentRequest{}
	// 参数校验
	if err := c.ShouldBindJSON(req); err != nil {
		response.BadRequest(c, "参数校验失败")
		return
	}
	// service层
	if error := service.CommentService.CreateComment(req, c); error != nil {
		response.InternalServerError(c, error.Error())
		return
	}

	response.Success(c, "success")
}

func (pt *CommentController) QueryComment(c *gin.Context) {
	// service层
	if error := service.CommentService.QueryComment(c); error != nil {
		response.InternalServerError(c, error.Error())
		return
	}
}
