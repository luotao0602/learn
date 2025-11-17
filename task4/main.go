package main

import (
	"task4/internal/config"
	"task4/pkg/db"
	"task4/pkg/log"
	"task4/router"
)

func main() {
	// 初始化配置
	config.InitConfig("configs/app.yaml")
	// 初始化日志
	logErr := log.Init()
	if logErr != nil {
		log.Logger.Info("日志初始化失败")
	}
	log.Logger.Info("配置初始化成功")
	log.Logger.Info("日志初始化成功")
	// 初始化DB
	db.InitDB()
	// 创建表
	db.CreateTable()
	log.Logger.Info("DB初始化成功")
	// 初始化路由
	router.StartService()
}
