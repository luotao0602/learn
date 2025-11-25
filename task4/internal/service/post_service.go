package service

import (
	"strconv"
	"task4/internal/dto"
	"task4/internal/model"
	"task4/pkg/db"
	"task4/pkg/exception"
	"task4/pkg/response"

	"github.com/gin-gonic/gin"
)

// 定义结构体
type postService struct{}

// 定义全局变量，单例
var PostService = new(postService)

// 创建文章
func (pt *postService) CreatePost(req *dto.PostRequest, c *gin.Context) error {
	// 获取上下文的user信息
	userId, exsit := c.Get("user_id")
	if !exsit {
		return exception.NewSystemException("User not authenticated")
	}

	db := db.GetGormDB()

	post := model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userId.(uint),
	}
	if error := db.Debug().Create(&post).Error; error != nil {
		return exception.NewSystemException(error.Error())
	}
	return nil
}

// 获取文章列表
func (pt *postService) QueryPostList(c *gin.Context) error {
	// 获取总数
	// 获取分页数
	pageNo, _ := strconv.Atoi(c.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	var total int64
	db := db.GetGormDB()
	db.Model(&model.Post{}).Count(&total)
	// 为0直接返回
	if total == 0 {
		response.Success(c, gin.H{
			"posts":     []model.Post{},
			"total":     total,
			"page_no":   pageNo,
			"page_size": pageSize,
		})
	}

	var posts []model.Post
	offset := (pageNo - 1) * pageSize

	if error := db.Debug().Preload("User").
		Order("created_at desc").
		Limit(pageSize).
		Offset(offset).
		Find(&posts).Error; error != nil {
		return exception.NewSystemException(error.Error())
	}

	response.Success(c, gin.H{
		"posts":     posts,
		"total":     total,
		"page_no":   pageNo,
		"page_size": pageSize,
	})
	return nil
}

// 获取文章详情
func (pt *postService) QueryPostDetail(c *gin.Context) error {
	var post model.Post
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, 400, err.Error())
	}

	if error := db.GetGormDB().Debug().Preload("User").
		Where("id = ?", id).
		Find(&post).Error; error != nil {
		return exception.NewSystemException(error.Error())
	}

	response.Success(c, post)
	return nil
}

// 更新文章
func (pt *postService) UpdatePost(req *dto.PostRequest, c *gin.Context) error {
	// 获取上下文的user信息
	userId, exsit := c.Get("user_id")
	if !exsit {
		return exception.NewSystemException("User not authenticated")
	}
	db := db.GetGormDB()
	var post model.Post
	if err := db.First(&post, req.ID).Error; err != nil {
		return exception.NewSystemException("Post not found")

	}
	// 检查是否是文章作者
	if post.UserID != userId.(uint) {
		return exception.NewSystemException("You can only update your own posts")
	}

	post.Title = req.Title
	post.Content = req.Content
	if error := db.Debug().Save(&post).Error; error != nil {
		return exception.NewSystemException(error.Error())
	}
	return nil
}
