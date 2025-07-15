package repository

import (
	"blog/internal/domain/entity"
)

// PostRepository 文章仓储接口
type PostRepository interface {
	// Create 创建文章
	Create(post *entity.Post) error
	// GetByID 根据ID获取文章
	GetByID(id uint) (*entity.Post, error)
	// GetAll 获取所有文章
	GetAll() ([]*entity.Post, error)
	// Update 更新文章
	Update(post *entity.Post) error
	// Delete 删除文章
	Delete(id uint) error
	// GetByCategory 根据分类获取文章
	GetByCategory(categoryID uint) ([]*entity.Post, error)
}
