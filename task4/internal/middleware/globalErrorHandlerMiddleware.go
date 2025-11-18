package middleware

import (
	"task4/pkg/log"

	"github.com/gin-gonic/gin"
)

// 自定义全局异常处理中间件，加入到 r.Use()中
func GlobleErrorHandlerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				switch {
				default:
					log.Logger.Error("e" + err.Error())
					return
				}
			}
		}
	}
}
