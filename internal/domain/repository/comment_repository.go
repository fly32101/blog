package repository

import (
	"blog/internal/domain/entity"
)

// CommentRepository 评论仓储接口
type CommentRepository interface {
	// Create 创建评论
	Create(comment *entity.Comment) error
	// GetByID 根据ID获取评论
	GetByID(id uint) (*entity.Comment, error)
	// GetByPostID 根据文章ID获取评论
	GetByPostID(postID uint) ([]*entity.Comment, error)
	// Delete 删除评论
	Delete(id uint) error
}
