package router

import (
	"task4/internal/middleware"
	"task4/pkg/log"

	"github.com/gin-gonic/gin"
)

func StartService() *gin.Engine {
	InitRouter()

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	err := r.Run(":8080")
	if err != nil {
		log.Logger.Error("service start failed")
	}
	log.Logger.Info("service start success")
}

func InitRouter() {
	gin.SetMode(gin.DebugMode)
	r := gin.New() // 创建纯净的 Gin 引擎
	/**
		gin.New() 会创建一个无默认中间件的纯净引擎实例，仅包含核心路由功能。
		对比 gin.Default()：gin.Default() = gin.New() + 2 个默认中间件（Logger 日志中间件
		 + Recovery 崩溃恢复中间件）。
	     适用场景：需要自定义中间件组合时（比如替换日志格式、添加权限校验中间件等）
	*/
	r.Use(gin.Recovery()) // 核心的崩溃恢复中间件
	// 新增自定义中间件
	r.Use(middleware.GlobleErrorHandlerMiddleWare(), middleware.LoggerMiddleWare())

}
