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
	TitleURL  string    `json:"title_url"`
	ViewCount uint      `json:"view_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewPost 创建新文章
func NewPost(title, content, author, titleURL string) *Post {
	now := time.Now()
	return &Post{
		Title:     title,
		Content:   content,
		Author:    author,
		TitleURL:  titleURL,
		ViewCount: 0,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Update 更新文章
func (p *Post) Update(title, content, titleURL string) {
	p.Title = title
	p.Content = content
	p.TitleURL = titleURL
	p.UpdatedAt = time.Now()
}
