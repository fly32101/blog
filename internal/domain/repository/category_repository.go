package repository

import (
	"blog/internal/domain/entity"
)

// CategoryRepository 分类仓储接口
type CategoryRepository interface {
	// Create 创建分类
	Create(category *entity.Category) error
	// GetByID 根据ID获取分类
	GetByID(id uint) (*entity.Category, error)
	// GetAll 获取所有分类
	GetAll() ([]*entity.Category, error)
	// Update 更新分类
	Update(category *entity.Category) error
	// Delete 删除分类
	Delete(id uint) error
	// AddPostToCategory 将文章添加到分类
	AddPostToCategory(postID, categoryID uint) error
	// RemovePostFromCategory 从分类中移除文章
	RemovePostFromCategory(postID, categoryID uint) error
	// GetCategoriesByPostID 获取文章的所有分类
	GetCategoriesByPostID(postID uint) ([]*entity.Category, error)
}
