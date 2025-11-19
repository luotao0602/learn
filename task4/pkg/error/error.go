package error

import "net/http"

// 定义系统错误的结构体
type AppError struct {
	Code    int    `json:"code"`
	ErrCode string `json:"error_code"`
	Message string `json:"message"`
}

func (err *AppError) Error() string {
	return err.Message
}

var (
	ErrSystem        = &AppError{Code: http.StatusInternalServerError, ErrCode: "SYSTEM_ERROR", Message: "系统异常,请稍后再试"}
	ErrInvalidParams = &AppError{http.StatusBadRequest, "INVALID_REQUEST", "请求参数错误"}
)
