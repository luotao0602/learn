package response

import (
	"net/http"
	error2 "task4/pkg/error"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

func Error(c *gin.Context, error *error2.AppError) {
	c.JSON(http.StatusOK, Response{
		Code:    error.Code,
		Message: error.Message,
		Data:    nil,
	})
}
