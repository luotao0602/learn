package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义全局异常处理中间件，加入到 r.Use()中
func LoggerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		// 跳至下一个中间件执行，完后再回溯处理c.Next()之后的业务代码
		c.Next()
		dealTime := time.Since(startTime)
		fmt.Sprintf("请求耗费时长: %v", dealTime)
	}
}
