package dto

type CommentRequest struct {
	ID      uint   `json:"id"`
	PostID  uint   `josn:"postId"  binding:"required,min=1"`
	Content string `json:"content" binding:"required,min=1,max=1000"`
}
