package entity

import (
	"time"
)

// Comment 评论实体
type Comment struct {
	ID        uint      `json:"id"`
	PostID    uint      `json:"post_id"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

// NewComment 创建新评论
func NewComment(postID uint, content, author string) *Comment {
	return &Comment{
		PostID:    postID,
		Content:   content,
		Author:    author,
		CreatedAt: time.Now(),
	}
}
