package service_test

import (
	"errors"
	"testing"
	"time"

	"blog/internal/domain/entity"
	"blog/internal/domain/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 模拟文章仓储
type MockPostRepository struct {
	mock.Mock
}

func (m *MockPostRepository) Create(post *entity.Post) error {
	args := m.Called(post)
	post.ID = 1 // 模拟数据库自增ID
	return args.Error(0)
}

func (m *MockPostRepository) GetByID(id uint) (*entity.Post, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Post), args.Error(1)
}

func (m *MockPostRepository) GetAll() ([]*entity.Post, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Post), args.Error(1)
}

func (m *MockPostRepository) Update(post *entity.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

func (m *MockPostRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPostRepository) GetByCategory(categoryID uint) ([]*entity.Post, error) {
	args := m.Called(categoryID)
	return args.Get(0).([]*entity.Post), args.Error(1)
}

// 模拟分类仓储
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Create(category *entity.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) GetByID(id uint) (*entity.Category, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Category), args.Error(1)
}

func (m *MockCategoryRepository) GetAll() ([]*entity.Category, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Category), args.Error(1)
}

func (m *MockCategoryRepository) Update(category *entity.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCategoryRepository) AddPostToCategory(postID, categoryID uint) error {
	args := m.Called(postID, categoryID)
	return args.Error(0)
}

func (m *MockCategoryRepository) RemovePostFromCategory(postID, categoryID uint) error {
	args := m.Called(postID, categoryID)
	return args.Error(0)
}

func (m *MockCategoryRepository) GetCategoriesByPostID(postID uint) ([]*entity.Category, error) {
	args := m.Called(postID)
	return args.Get(0).([]*entity.Category), args.Error(1)
}

// 测试创建文章
func TestCreatePost(t *testing.T) {
	mockPostRepo := new(MockPostRepository)
	mockCategoryRepo := new(MockCategoryRepository)

	title := "测试标题"
	content := "测试内容"
	author := "测试作者"

	mockPostRepo.On("Create", mock.AnythingOfType("*entity.Post")).Return(nil)

	postService := service.NewPostService(mockPostRepo, mockCategoryRepo)
	post, err := postService.CreatePost(title, content, author)

	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, title, post.Title)
	assert.Equal(t, content, post.Content)
	assert.Equal(t, author, post.Author)
	mockPostRepo.AssertExpectations(t)
}

// 测试获取文章
func TestGetPostByID(t *testing.T) {
	mockPostRepo := new(MockPostRepository)
	mockCategoryRepo := new(MockCategoryRepository)

	now := time.Now()
	expectedPost := &entity.Post{
		ID:        1,
		Title:     "测试标题",
		Content:   "测试内容",
		Author:    "测试作者",
		CreatedAt: now,
		UpdatedAt: now,
	}

	mockPostRepo.On("GetByID", uint(1)).Return(expectedPost, nil)

	postService := service.NewPostService(mockPostRepo, mockCategoryRepo)
	post, err := postService.GetPostByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedPost, post)
	mockPostRepo.AssertExpectations(t)
}

// 测试获取不存在的文章
func TestGetPostByIDNotFound(t *testing.T) {
	mockPostRepo := new(MockPostRepository)
	mockCategoryRepo := new(MockCategoryRepository)

	expectedError := errors.New("文章不存在")
	mockPostRepo.On("GetByID", uint(999)).Return((*entity.Post)(nil), expectedError)

	postService := service.NewPostService(mockPostRepo, mockCategoryRepo)
	post, err := postService.GetPostByID(999)

	assert.Error(t, err)
	assert.Nil(t, post)
	assert.Equal(t, expectedError, err)
	mockPostRepo.AssertExpectations(t)
}

// 测试更新文章
func TestUpdatePost(t *testing.T) {
	mockPostRepo := new(MockPostRepository)
	mockCategoryRepo := new(MockCategoryRepository)

	now := time.Now()
	existingPost := &entity.Post{
		ID:        1,
		Title:     "旧标题",
		Content:   "旧内容",
		Author:    "测试作者",
		CreatedAt: now,
		UpdatedAt: now,
	}

	newTitle := "新标题"
	newContent := "新内容"

	mockPostRepo.On("GetByID", uint(1)).Return(existingPost, nil)
	mockPostRepo.On("Update", mock.AnythingOfType("*entity.Post")).Return(nil)

	postService := service.NewPostService(mockPostRepo, mockCategoryRepo)
	post, err := postService.UpdatePost(1, newTitle, newContent)

	assert.NoError(t, err)
	assert.Equal(t, newTitle, post.Title)
	assert.Equal(t, newContent, post.Content)
	mockPostRepo.AssertExpectations(t)
}

// 测试删除文章
func TestDeletePost(t *testing.T) {
	mockPostRepo := new(MockPostRepository)
	mockCategoryRepo := new(MockCategoryRepository)

	mockPostRepo.On("Delete", uint(1)).Return(nil)

	postService := service.NewPostService(mockPostRepo, mockCategoryRepo)
	err := postService.DeletePost(1)

	assert.NoError(t, err)
	mockPostRepo.AssertExpectations(t)
}
