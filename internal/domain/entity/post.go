package entity

import (
	"time"
)

// Post 博客文章实体
type Post struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewPost 创建新文章
func NewPost(title, content, author string) *Post {
	now := time.Now()
	return &Post{
		Title:     title,
		Content:   content,
		Author:    author,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Update 更新文章
func (p *Post) Update(title, content string) {
	p.Title = title
	p.Content = content
	p.UpdatedAt = time.Now()
}
