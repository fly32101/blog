package application

import (
	"blog/internal/domain/entity"
	"blog/internal/domain/service"
	"blog/internal/interfaces/dto"
)

// CategoryApp 分类应用服务
type CategoryApp struct {
	categoryService *service.CategoryService
}

// NewCategoryApp 创建分类应用服务
func NewCategoryApp(categoryService *service.CategoryService) *CategoryApp {
	return &CategoryApp{
		categoryService: categoryService,
	}
}

// CreateCategory 创建分类
func (a *CategoryApp) CreateCategory(req *dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	category, err := a.categoryService.CreateCategory(req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	return convertToCategoryResponse(category), nil
}

// GetCategoryByID 根据ID获取分类
func (a *CategoryApp) GetCategoryByID(id uint) (*dto.CategoryResponse, error) {
	category, err := a.categoryService.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return convertToCategoryResponse(category), nil
}

// GetAllCategories 获取所有分类
func (a *CategoryApp) GetAllCategories() ([]*dto.CategoryResponse, error) {
	categories, err := a.categoryService.GetAllCategories()
	if err != nil {
		return nil, err
	}

	return convertToCategoryResponses(categories), nil
}

// UpdateCategory 更新分类
func (a *CategoryApp) UpdateCategory(id uint, req *dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	category, err := a.categoryService.UpdateCategory(id, req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	return convertToCategoryResponse(category), nil
}

// DeleteCategory 删除分类
func (a *CategoryApp) DeleteCategory(id uint) error {
	return a.categoryService.DeleteCategory(id)
}

// 转换为分类响应
func convertToCategoryResponse(category *entity.Category) *dto.CategoryResponse {
	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}
