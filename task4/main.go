package main

import (
	"task4/internal/config"
	"task4/pkg/db"
	"task4/pkg/log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config.InitConfig("configs/app.yaml")
	// 初始化日志
	logErr := log.Init()
	if logErr != nil {
		panic(logErr)
	}
	log.Logger.Info("配置初始化成功")
	log.Logger.Info("日志初始化成功")
	// 初始化DB
	db.InitDB()
	log.Logger.Info("DB初始化成功")
	// 初始化路由

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
