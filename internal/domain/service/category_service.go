package service

import (
	"blog/internal/domain/entity"
	"blog/internal/domain/repository"
)

// CategoryService 分类领域服务
type CategoryService struct {
	categoryRepo repository.CategoryRepository
}

// NewCategoryService 创建分类服务
func NewCategoryService(categoryRepo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

// CreateCategory 创建分类
func (s *CategoryService) CreateCategory(name, description string) (*entity.Category, error) {
	category := entity.NewCategory(name, description)
	err := s.categoryRepo.Create(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// GetCategoryByID 根据ID获取分类
func (s *CategoryService) GetCategoryByID(id uint) (*entity.Category, error) {
	return s.categoryRepo.GetByID(id)
}

// GetAllCategories 获取所有分类
func (s *CategoryService) GetAllCategories() ([]*entity.Category, error) {
	return s.categoryRepo.GetAll()
}

// UpdateCategory 更新分类
func (s *CategoryService) UpdateCategory(id uint, name, description string) (*entity.Category, error) {
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	category.Name = name
	category.Description = description

	err = s.categoryRepo.Update(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// DeleteCategory 删除分类
func (s *CategoryService) DeleteCategory(id uint) error {
	return s.categoryRepo.Delete(id)
}

// GetCategoriesByPostID 获取文章的所有分类
func (s *CategoryService) GetCategoriesByPostID(postID uint) ([]*entity.Category, error) {
	return s.categoryRepo.GetCategoriesByPostID(postID)
}
