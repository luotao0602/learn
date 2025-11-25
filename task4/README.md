一个使用Go语言、Gin框架和GORM库开发的个人博客系统后端API。

功能特性
用户注册和登录
JWT认证和授权
文章的CRUD操作
评论功能
统一的错误处理
日志记录
数据库自动迁移
技术栈

Go 1.23.0
Gin - Web框架
GORM - ORM库
MYSQL - 数据库
JWT - 身份认证
Logrus - 日志库
bcrypt - 密码加密
项目结构
golang_blog/
├── main.go          # 程序入口
├── configs/
│   └── app.yaml      # 数据库相关yaml文件
├── internal/
│   └── config      
│       └── config.go               # 数据库初始化配置
│   └── controller                 
│       └── auth_controller.go         # 认证控制器
│       └── comment_controller.go      # 评论控制器
│       └── post_controller.go         # 文章控制器
│       └── user_controller.go         # 用户控制器
│   └── dto                            # 请求体和响应结构体目录
│       └── auth_response.go           # 认证相应结构体
│       └── comment_request.go         # 评论请求结构体
│       └── post_request.go            # 文章请求结构体
│       └── register_request.go        # 用户注册请求结构体
│   └── middleware      # 中间件相关
│       └── auth_middleware.go         # 认证中间件
│       └── global_error_handler_middleware.go        # 全局异常中间件
│       └── logger_middleware.go       # 日志中间件
│   └── model      # 数据库表对应的结构体
│       └── comment.go                 # 评论表
│       └── post.go                 
│       └── user.go                 
│   └── service                        # 业务层
│       └── comment_service.go         # 评论的业务层
│       └── post_service.go                 
│       └── user_service.go                 
├── pkg/                               #公共目录
│   └── db                             
│       └── db.go                      # db连接
│   └── exception                       
│       └── exception.go               # 自定义异常
│   └── response                        
│       └── response.go                # 响应
│   └── utils                          #工具类相关  
│       └── jwt.go                     # token生成和解密 
├── router/
│   └── router.go                     # 路由配置
├── go.mod
├── go.sum
└── README.md

数据库设计
Users表
id (主键)
username (用户名，唯一)
email (邮箱，唯一)
password (加密密码)
created_at, updated_at, deleted_at

Posts表
id (主键)
title (标题)
content (内容)
user_id (外键，关联users表)
created_at, updated_at, deleted_at


Comments表
id (主键)
content (内容)
user_id (外键，关联users表)
post_id (外键，关联posts表)
created_at, updated_at, deleted_at
