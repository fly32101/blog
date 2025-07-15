package dto

import (
	"time"
)

// CreateCommentRequest 创建评论请求
type CreateCommentRequest struct {
	PostID  uint   `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
	Author  string `json:"author" binding:"required"`
}

// CommentResponse 评论响应
type CommentResponse struct {
	ID        uint      `json:"id"`
	PostID    uint      `json:"post_id"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}
