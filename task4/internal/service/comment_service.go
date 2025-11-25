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
type commentService struct{}

// 定义全局变量，单例
var CommentService = new(commentService)

// 创建文章
func (pt *commentService) CreateComment(req *dto.CommentRequest, c *gin.Context) error {
	// 获取上下文的user信息
	userId, exsit := c.Get("user_id")
	if !exsit {
		return exception.NewSystemException("User not authenticated")
	}

	db := db.GetGormDB()

	comment := model.Comment{
		Content: req.Content,
		UserID:  userId.(uint),
		PostID:  req.PostID,
	}
	if error := db.Debug().Create(&comment).Error; error != nil {
		return exception.NewSystemException(error.Error())
	}
	return nil
}

// 获取文章详情
func (pt *commentService) QueryComment(c *gin.Context) error {
	var comments []model.Comment
	pageNo, _ := strconv.Atoi(c.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	postId, _ := strconv.Atoi(c.Param("post_id"))
	var total int64
	if err := db.GetGormDB().Model(&model.Comment{}).Where("post_id = ? ", postId).Count(&total).Error; err != nil {
		return exception.NewSystemException(err.Error())
	}

	offSize := (pageNo - 1) * pageSize
	if err := db.GetGormDB().Preload("User").Where("post_id =? ", postId).
		Order("created_at desc").
		Offset(offSize).
		Find(&comments).Error; err != nil {
		return exception.NewSystemException(err.Error())
	}

	response.Success(c, gin.H{"comments": comments,
		"total":     total,
		"page_no":   pageNo,
		"page_size": pageSize})
	return nil
}
