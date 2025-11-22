package exception

import "net/http"

// 定义系统异常的结构体
type Exception struct {
	Code    int    `json:"code"`
	ErrCode string `json:"error_code"`
	Message string `json:"message"`
}

func (err *Exception) Error() string {
	return err.Message
}

var (
	ErrSystem        = &Exception{Code: http.StatusInternalServerError, ErrCode: "SYSTEM_ERROR", Message: "系统异常,请稍后再试"}
	ErrInvalidParams = &Exception{http.StatusBadRequest, "INVALID_REQUEST", "请求参数错误"}
)

func NewSystemException(msg string) *Exception {
	return &Exception{
		Message: msg,
	}
}
