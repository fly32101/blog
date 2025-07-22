package dto

import (
	"time"
)

// CreatePostRequest 创建文章请求
type CreatePostRequest struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Author      string `json:"author" binding:"required"`
	CategoryIDs []uint `json:"category_ids"`
}

// UpdatePostRequest 更新文章请求
type UpdatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// PostResponse 文章响应
type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	ViewCount uint      `json:"view_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PostDetailResponse 文章详情响应
type PostDetailResponse struct {
	ID         uint                `json:"id"`
	Title      string              `json:"title"`
	Content    string              `json:"content"`
	Author     string              `json:"author"`
	ViewCount  uint                `json:"view_count"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	Categories []*CategoryResponse `json:"categories"`
}
