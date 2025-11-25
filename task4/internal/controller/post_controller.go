package controller

import (
	"task4/internal/dto"
	"task4/internal/service"
	"task4/pkg/response"

	"github.com/gin-gonic/gin"
)

// 定义请求
type PostController struct{}

// var AuthController = new(authController)
func (pt *PostController) CreatePost(c *gin.Context) {
	req := &dto.PostRequest{}
	// 参数校验
	if err := c.ShouldBindJSON(req); err != nil {
		response.BadRequest(c, "参数校验失败")
		return
	}
	// service层
	if error := service.PostService.CreatePost(req, c); error != nil {
		response.InternalServerError(c, error.Error())
		return
	}

	response.Success(c, "success")
}

func (pt *PostController) QueryPostList(c *gin.Context) {
	// service层
	if error := service.PostService.QueryPostList(c); error != nil {
		response.InternalServerError(c, error.Error())
		return
	}
}

func (pt *PostController) QueryPostDetail(c *gin.Context) {
	// service层
	if error := service.PostService.QueryPostDetail(c); error != nil {
		response.InternalServerError(c, error.Error())
		return
	}
}

func (pt *PostController) UpdatePost(c *gin.Context) {
	req := &dto.PostRequest{}
	// 参数校验
	if err := c.ShouldBindJSON(req); err != nil {
		response.BadRequest(c, "参数校验失败")
		return
	}
	// service层
	if error := service.PostService.UpdatePost(req, c); error != nil {
		response.InternalServerError(c, error.Error())
		return
	}

	response.Success(c, "success")
}
