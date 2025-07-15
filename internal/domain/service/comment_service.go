package service

import (
	"blog/internal/domain/entity"
	"blog/internal/domain/repository"
)

// CommentService 评论领域服务
type CommentService struct {
	commentRepo repository.CommentRepository
	postRepo    repository.PostRepository
}

// NewCommentService 创建评论服务
func NewCommentService(commentRepo repository.CommentRepository, postRepo repository.PostRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

// CreateComment 创建评论
func (s *CommentService) CreateComment(postID uint, content, author string) (*entity.Comment, error) {
	// 验证文章是否存在
	_, err := s.postRepo.GetByID(postID)
	if err != nil {
		return nil, err
	}

	comment := entity.NewComment(postID, content, author)
	err = s.commentRepo.Create(comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// GetCommentByID 根据ID获取评论
func (s *CommentService) GetCommentByID(id uint) (*entity.Comment, error) {
	return s.commentRepo.GetByID(id)
}

// GetCommentsByPostID 根据文章ID获取评论
func (s *CommentService) GetCommentsByPostID(postID uint) ([]*entity.Comment, error) {
	return s.commentRepo.GetByPostID(postID)
}

// DeleteComment 删除评论
func (s *CommentService) DeleteComment(id uint) error {
	return s.commentRepo.Delete(id)
}
