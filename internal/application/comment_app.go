package application

import (
	"blog/internal/domain/entity"
	"blog/internal/domain/service"
	"blog/internal/interfaces/dto"
)

// CommentApp 评论应用服务
type CommentApp struct {
	commentService *service.CommentService
}

// NewCommentApp 创建评论应用服务
func NewCommentApp(commentService *service.CommentService) *CommentApp {
	return &CommentApp{
		commentService: commentService,
	}
}

// CreateComment 创建评论
func (a *CommentApp) CreateComment(req *dto.CreateCommentRequest) (*dto.CommentResponse, error) {
	comment, err := a.commentService.CreateComment(req.PostID, req.Content, req.Author)
	if err != nil {
		return nil, err
	}

	return convertToCommentResponse(comment), nil
}

// GetCommentByID 根据ID获取评论
func (a *CommentApp) GetCommentByID(id uint) (*dto.CommentResponse, error) {
	comment, err := a.commentService.GetCommentByID(id)
	if err != nil {
		return nil, err
	}

	return convertToCommentResponse(comment), nil
}

// GetCommentsByPostID 根据文章ID获取评论
func (a *CommentApp) GetCommentsByPostID(postID uint) ([]*dto.CommentResponse, error) {
	comments, err := a.commentService.GetCommentsByPostID(postID)
	if err != nil {
		return nil, err
	}

	var commentResponses []*dto.CommentResponse
	for _, comment := range comments {
		commentResponses = append(commentResponses, convertToCommentResponse(comment))
	}

	return commentResponses, nil
}

// DeleteComment 删除评论
func (a *CommentApp) DeleteComment(id uint) error {
	return a.commentService.DeleteComment(id)
}

// 转换为评论响应
func convertToCommentResponse(comment *entity.Comment) *dto.CommentResponse {
	return &dto.CommentResponse{
		ID:        comment.ID,
		PostID:    comment.PostID,
		Content:   comment.Content,
		Author:    comment.Author,
		CreatedAt: comment.CreatedAt,
	}
}
