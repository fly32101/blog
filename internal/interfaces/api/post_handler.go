package api

import (
	"net/http"
	"strconv"

	"blog/internal/application"
	"blog/internal/interfaces/dto"
	"blog/pkg/utils"

	"github.com/gin-gonic/gin"
)

// PostHandler 文章处理器
type PostHandler struct {
	postApp *application.PostApp
}

// NewPostHandler 创建文章处理器
func NewPostHandler(postApp *application.PostApp) *PostHandler {
	return &PostHandler{
		postApp: postApp,
	}
}

// Register 注册路由
func (h *PostHandler) Register(router *gin.RouterGroup) {
	posts := router.Group("/posts")
	{
		posts.POST("", h.CreatePost)
		posts.GET("", h.GetAllPosts)
		posts.GET("/:id", h.GetPostByID)
		posts.PUT("/:id", h.UpdatePost)
		posts.DELETE("/:id", h.DeletePost)
		posts.GET("/category/:id", h.GetPostsByCategory)
	}
}

// CreatePost 创建文章
func (h *PostHandler) CreatePost(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "请求参数错误"))
		return
	}

	post, err := h.postApp.CreatePost(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err, "创建文章失败"))
		return
	}

	c.JSON(http.StatusCreated, utils.NewSuccessResponse(post, "文章创建成功"))
}

// GetPostByID 根据ID获取文章
func (h *PostHandler) GetPostByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "无效的ID"))
		return
	}

	post, err := h.postApp.GetPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NewErrorResponse(err, "文章不存在"))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(post, "获取文章详情成功"))
}

// GetAllPosts 获取所有文章
func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.postApp.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err, "获取文章列表失败"))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(posts, "获取文章列表成功"))
}

// UpdatePost 更新文章
func (h *PostHandler) UpdatePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "无效的ID"))
		return
	}

	var req dto.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "请求参数错误"))
		return
	}

	post, err := h.postApp.UpdatePost(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err, "更新文章失败"))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(post, "文章更新成功"))
}

// DeletePost 删除文章
func (h *PostHandler) DeletePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "无效的ID"))
		return
	}

	err = h.postApp.DeletePost(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err, "删除文章失败"))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(nil, "文章删除成功"))
}

// GetPostsByCategory 根据分类获取文章
func (h *PostHandler) GetPostsByCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "无效的分类ID"))
		return
	}

	posts, err := h.postApp.GetPostsByCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err, "获取分类文章失败"))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(posts, "获取分类文章成功"))
}
