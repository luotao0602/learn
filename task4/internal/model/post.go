package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系  Preload 预加载  foreignkey关联外键
	User    User      `json:"user,omitempty" gorm:"foreignkey:UserID"`
	Comment []Comment `json:"comments,omitempty" gorm:"foreignkey:PostID"`
}
