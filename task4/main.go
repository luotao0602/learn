package main

import (
	"task4/internal/config"
	"task4/pkg/db"
	"task4/router"

	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化日志
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("日志初始化成功")
	// 初始化配置
	config.InitConfig("configs/app.yaml")
	logrus.Info("初始化配置成功")

	// 初始化DB
	db.InitDB()
	// 创建表
	// db.CreateTable()
	logrus.Info("DB初始化成功")
	// 初始化路由
	r := router.InitRouter()

	//启动服务
	err := r.Run(":8080")
	if err != nil {
		logrus.Error("service start failed")
	}
	logrus.Info("service start success")
}
