package application

import (
	"blog/internal/domain/entity"
	"blog/internal/domain/service"
	"blog/internal/interfaces/dto"
)

// PostApp 文章应用服务
type PostApp struct {
	postService     *service.PostService
	categoryService *service.CategoryService
}

// NewPostApp 创建文章应用服务
func NewPostApp(postService *service.PostService, categoryService *service.CategoryService) *PostApp {
	return &PostApp{
		postService:     postService,
		categoryService: categoryService,
	}
}

// CreatePost 创建文章
func (a *PostApp) CreatePost(req *dto.CreatePostRequest) (*dto.PostResponse, error) {
	post, err := a.postService.CreatePost(req.Title, req.Content, req.Author)
	if err != nil {
		return nil, err
	}

	// 如果有分类，添加文章到分类
	if len(req.CategoryIDs) > 0 {
		for _, categoryID := range req.CategoryIDs {
			err = a.postService.AddPostToCategory(post.ID, categoryID)
			if err != nil {
				return nil, err
			}
		}
	}

	return convertToPostResponse(post), nil
}

// GetPostByID 根据ID获取文章
func (a *PostApp) GetPostByID(id uint) (*dto.PostDetailResponse, error) {
	post, err := a.postService.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	// 获取文章分类
	categories, err := a.categoryService.GetCategoriesByPostID(id)
	if err != nil {
		return nil, err
	}

	return &dto.PostDetailResponse{
		ID:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		Author:     post.Author,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
		Categories: convertToCategoryResponses(categories),
	}, nil
}

// GetAllPosts 获取所有文章
func (a *PostApp) GetAllPosts() ([]*dto.PostResponse, error) {
	posts, err := a.postService.GetAllPosts()
	if err != nil {
		return nil, err
	}

	var postResponses []*dto.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, convertToPostResponse(post))
	}

	return postResponses, nil
}

// UpdatePost 更新文章
func (a *PostApp) UpdatePost(id uint, req *dto.UpdatePostRequest) (*dto.PostResponse, error) {
	post, err := a.postService.UpdatePost(id, req.Title, req.Content)
	if err != nil {
		return nil, err
	}

	return convertToPostResponse(post), nil
}

// DeletePost 删除文章
func (a *PostApp) DeletePost(id uint) error {
	return a.postService.DeletePost(id)
}

// GetPostsByCategory 根据分类获取文章
func (a *PostApp) GetPostsByCategory(categoryID uint) ([]*dto.PostResponse, error) {
	posts, err := a.postService.GetPostsByCategory(categoryID)
	if err != nil {
		return nil, err
	}

	var postResponses []*dto.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, convertToPostResponse(post))
	}

	return postResponses, nil
}

// 转换为文章响应
func convertToPostResponse(post *entity.Post) *dto.PostResponse {
	return &dto.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Author:    post.Author,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

// 转换为分类响应列表
func convertToCategoryResponses(categories []*entity.Category) []*dto.CategoryResponse {
	var categoryResponses []*dto.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, &dto.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return categoryResponses
}
