package api

import (
	"net/http"
	"strconv"

	"blog/internal/application"
	"blog/internal/interfaces/dto"
	"blog/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CommentHandler 评论处理器
type CommentHandler struct {
	commentApp *application.CommentApp
}

// NewCommentHandler 创建评论处理器
func NewCommentHandler(commentApp *application.CommentApp) *CommentHandler {
	return &CommentHandler{
		commentApp: commentApp,
	}
}

// Register 注册路由
func (h *CommentHandler) Register(router *gin.RouterGroup) {
	comments := router.Group("/comments")
	{
		comments.POST("", h.CreateComment)
		comments.GET("/:id", h.GetCommentByID)
		comments.GET("/post/:id", h.GetCommentsByPostID)
		comments.DELETE("/:id", h.DeleteComment)
	}
}

// CreateComment 创建评论
func (h *CommentHandler) CreateComment(c *gin.Context) {
	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "请求参数错误"))
		return
	}

	comment, err := h.commentApp.CreateComment(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err, "创建评论失败"))
		return
	}

	c.JSON(http.StatusCreated, utils.NewSuccessResponse(comment, "评论创建成功"))
}

// GetCommentByID 根据ID获取评论
func (h *CommentHandler) GetCommentByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "无效的ID"))
		return
	}

	comment, err := h.commentApp.GetCommentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NewErrorResponse(err, "评论不存在"))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(comment, "获取评论成功"))
}

// GetCommentsByPostID 根据文章ID获取评论
func (h *CommentHandler) GetCommentsByPostID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "无效的文章ID"))
		return
	}

	comments, err := h.commentApp.GetCommentsByPostID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err, "获取文章评论失败"))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(comments, "获取文章评论成功"))
}

// DeleteComment 删除评论
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err, "无效的ID"))
		return
	}

	err = h.commentApp.DeleteComment(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err, "删除评论失败"))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(nil, "评论删除成功"))
}
