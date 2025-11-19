package router

import (
	"task4/internal/handler"
	"task4/internal/middleware"
	"task4/pkg/log"

	"github.com/gin-gonic/gin"
)

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
	// 路由

	apiV1 := r.Group("/api/v1")
	{
		auth := apiV1.Group("/auth")
		{
			// 这里的 handler.Register 是作为一个「函数值」（可以理解为 “函数的引用”）传递给 Gin，而非执行它。
			// 此时不需要传参，因为：
			// 函数还没被执行，只是告诉框架 “要执行哪个函数”；
			// 当请求真正到来时，Gin 会自动创建 *gin.Context 实例（封装了请求信息、响应工具等），
			// 然后调用 handler.Register(c)，把上下文参数 c 注入进去。
			auth.POST("/register/", handler.Register)
		}
	}

	err := r.Run(":8080")
	if err != nil {
		log.Logger.Error("service start failed")
	}
	log.Logger.Info("service start success")
}
