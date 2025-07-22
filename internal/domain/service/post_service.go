package service

import (
	"blog/internal/domain/entity"
	"blog/internal/domain/repository"
)

// PostService 文章领域服务
type PostService struct {
	postRepo     repository.PostRepository
	categoryRepo repository.CategoryRepository
}

// NewPostService 创建文章服务
func NewPostService(postRepo repository.PostRepository, categoryRepo repository.CategoryRepository) *PostService {
	return &PostService{
		postRepo:     postRepo,
		categoryRepo: categoryRepo,
	}
}

// CreatePost 创建文章
func (s *PostService) CreatePost(title, content, author, titleURL string) (*entity.Post, error) {
	post := entity.NewPost(title, content, author, titleURL)
	err := s.postRepo.Create(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// GetPostByID 根据ID获取文章
func (s *PostService) GetPostByID(id uint) (*entity.Post, error) {
	return s.postRepo.GetByID(id)
}

// GetAllPosts 获取所有文章
func (s *PostService) GetAllPosts() ([]*entity.Post, error) {
	return s.postRepo.GetAll()
}

// UpdatePost 更新文章
func (s *PostService) UpdatePost(id uint, title, content, titleURL string) (*entity.Post, error) {
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	post.Update(title, content, titleURL)
	err = s.postRepo.Update(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// IncrementViewCount 增加文章阅读量
func (s *PostService) IncrementViewCount(id uint) error {
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return err
	}

	post.ViewCount++
	err = s.postRepo.Update(post)
	if err != nil {
		return err
	}

	return nil
}

// DeletePost 删除文章
func (s *PostService) DeletePost(id uint) error {
	return s.postRepo.Delete(id)
}

// AddPostToCategory 将文章添加到分类
func (s *PostService) AddPostToCategory(postID, categoryID uint) error {
	return s.categoryRepo.AddPostToCategory(postID, categoryID)
}

// GetPostsByCategory 根据分类获取文章
func (s *PostService) GetPostsByCategory(categoryID uint) ([]*entity.Post, error) {
	return s.postRepo.GetByCategory(categoryID)
}
