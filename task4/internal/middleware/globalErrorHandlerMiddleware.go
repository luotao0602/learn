package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GlobleErrorHandlerMiddleWare 全局错误处理中间件
func GlobleErrorHandlerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 捕获 Panic 异常
		defer func() {
			if err := recover(); err != nil {
				// 记录堆栈信息
				stack := string(debug.Stack())
				logrus.WithFields(logrus.Fields{
					"error":   err,
					"path":    c.Request.URL.Path,
					"method":  c.Request.Method,
					"stack":   stack,
					"headers": c.Request.Header, // 可选：记录请求头，便于排查
				}).Error("Panic recovered")

				// 统一响应：包含错误码和详细信息
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "Internal server error",
					"detail":  fmt.Sprintf("%v", err), // 可选：返回错误详情（生产环境慎用）
				})
				c.Abort()
			}
		}()

		// 执行后续业务逻辑
		c.Next()

		// 处理普通错误（非 Panic）
		if len(c.Errors) > 0 {
			// 获取最后一个错误（通常是最关键的错误）
			lastErr := c.Errors.Last()

			// 记录错误日志（包含错误类型和堆栈）
			logrus.WithFields(logrus.Fields{
				"error":  lastErr.Err,
				"path":   c.Request.URL.Path,
				"method": c.Request.Method,
				"type":   lastErr.Type, // 错误类型（如 binding error、DB error）
				"stack":  string(debug.Stack()),
			}).Error("Request error")

			// 根据错误类型返回对应状态码
			statusCode := http.StatusInternalServerError
			switch lastErr.Type {
			case gin.ErrorTypeBind: // 参数绑定错误
				statusCode = http.StatusBadRequest
			case gin.ErrorTypeRender: // 响应渲染错误
				statusCode = http.StatusInternalServerError
				// 可扩展自定义错误类型
			}

			// 统一响应格式
			c.JSON(statusCode, gin.H{
				"code":    statusCode,
				"message": lastErr.Err.Error(),
			})
			c.Abort()
		}
	}
}
