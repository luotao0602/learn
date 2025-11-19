package handler

import (
	"task4/internal/model"
	error2 "task4/pkg/error"
	"task4/pkg/response"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	user := &model.User{}
	// 参数校验
	if err := c.ShouldBindJSON(user); err != nil {
		response.Error(c, error2.ErrInvalidParams)
	}
	// service层

	response.Success(c, "success", "success")
}
