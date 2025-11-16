package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
type User struct {
	gorm.Model
	ID       uint
	Name     string
	PostSize uint
	Posts    []Post
}
type Post struct {
	gorm.Model
	ID            uint
	Title         string
	Content       string
	UserID        uint
	Comments      []Comment
	CommentStatus string
}
type Comment struct {
	gorm.Model // 内置字段：ID（主键）、CreatedAt（创建时间）、UpdatedAt（更新时间）、DeletedAt（软删除标记）
	ID         uint
	Content    string
	PostId     uint
}

var db *gorm.DB

func initDb() {
	const url = "root:123456@tcp(127.0.0.1:3306)/lt?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println("connect db failed,err: %v", err)
		panic("connect DB failed")
	}
}

// 创建表
func createTable() {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

// 初始化表数据
func initTableData() {
	user := User{Name: "张三"}
	post1 := Post{Title: "第一篇文章", Content: "内容A", CommentStatus: "有评论"}
	post2 := Post{Title: "第二篇文章", Content: "内容B", CommentStatus: "有评论"}
	comment1 := Comment{Content: "评论1"}
	comment2 := Comment{Content: "评论2"}
	comment3 := Comment{Content: "评论3"}

	post1.Comments = append(post1.Comments, comment1, comment2)
	post2.Comments = append(post2.Comments, comment3)
	user.Posts = append(user.Posts, post1, post2)
	result := db.Debug().Create(&user)
	if result.Error != nil {
		log.Fatalf("创建用户失败: %v", result.Error)
	}
}

func main() {
	initDb()
	// createTable()
	// initTableData()

	// 使用Gorm查询某个用户发布的所有文章及其对应的评论信息
	// fmt.Println(getUserAllInfoByID(1))

	// 使用Gorm查询评论数量最多的文章信息
	// fmt.Println(getMaxCommentOfPosts())

	//创建post
	//createPost()

	//删除Comment
	deleteComment()
}

func deleteComment() {
	result := db.Debug().Delete(&Comment{ID: 6, PostId: 4})
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		fmt.Errorf("delete error, detail: %v", result.Error)
	}
	fmt.Println("delete Comment success")
}

func createPost() {
	post1 := Post{Title: "第1篇文章", Content: "内容A", CommentStatus: "有评论", UserID: 1}
	result := db.Debug().Create(&post1)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		fmt.Errorf("数据插入失败，post_id: %v", post1.ID)
	}
}

/*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/

func getUserAllInfoByID(userID uint) User {
	var user User
	db.Debug().
		Preload("Posts").
		Preload("Posts.Comments").
		First(&user, userID)
	return user
}

func getMaxCommentOfPosts() Post {
	var post Post
	most := db.Debug().Model(&Comment{}).
		Select("post_id,count(*) as total_comment").
		Group("post_id").
		Order("total_comment desc").
		Limit(1)
	//关联子查询，查询出文章信息
	db.Debug().
		Model(&Post{}).
		Select("post.*").
		Joins("join (?) most on post.id=most.post_id", most).
		First(&post)

	db.Debug().Model(&Comment{PostId: post.ID}).Find(&post.Comments)
	return post
}

/*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func (p *Post) AfterCreate(db *gorm.DB) error {
	var count int64
	err := db.Debug().Model(&Post{UserID: p.UserID}).Count(&count).Error
	if err != nil {
		return err
	}

	result := db.Debug().Model(&User{ID: p.UserID}).Update("post_size", count)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("更新失败")
	}
	return nil
}

// Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论
func (c *Comment) AfterDelete(db *gorm.DB) error {
	var count int64
	err := db.Debug().Model(&Comment{}).Where("post_id = ? and deleted_at IS NOT NULL", c.PostId).Count(&count).Error
	if err != nil {
		return err
	}
	if count != 0 {
		return nil
	}

	result := db.Debug().Model(&Post{ID: c.PostId}).Update("comment_status", "无评论")
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("更新失败")
	}
	return nil
}
