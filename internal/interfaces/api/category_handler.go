package api

import (
	"net/http"
	"strconv"

	"blog/internal/application"
	"blog/internal/interfaces/dto"
	"blog/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CategoryHandler 分类处理器
type CategoryHandler struct {
	categoryApp *application.CategoryApp
}

// NewCategoryHandler 创建分类处理器
func NewCategoryHandler(categoryApp *application.CategoryApp) *CategoryHandler {
	return &CategoryHandler{
		categoryApp: categoryApp,
	}
}

// Register 注册路由
func (h *CategoryHandler) Register(router *gin.RouterGroup) {
	categories := router.Group("/categories")
	{
		categories.POST("", h.CreateCategory)
		categories.GET("", h.GetAllCategories)
		categories.GET("/:id", h.GetCategoryByID)
		categories.PUT("/:id", h.UpdateCategory)
		categories.DELETE("/:id", h.DeleteCategory)
	}
}

// CreateCategory 创建分类
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "请求参数错误"))
		return
	}

	category, err := h.categoryApp.CreateCategory(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err, "创建分类失败"))
		return
	}

	c.JSON(http.StatusCreated, utils.NewSuccessResponse(category, "分类创建成功"))
}

// GetCategoryByID 根据ID获取分类
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "无效的ID"))
		return
	}

	category, err := h.categoryApp.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NewErrorResponse(err, "分类不存在"))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(category, "获取分类详情成功"))
}

// GetAllCategories 获取所有分类
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.categoryApp.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err, "获取分类列表失败"))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(categories, "获取分类列表成功"))
}

// UpdateCategory 更新分类
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "无效的ID"))
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "请求参数错误"))
		return
	}

	category, err := h.categoryApp.UpdateCategory(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err, "更新分类失败"))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(category, "分类更新成功"))
}

// DeleteCategory 删除分类
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "无效的ID"))
		return
	}

	err = h.categoryApp.DeleteCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err, "删除分类失败"))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(nil, "分类删除成功"))
}
