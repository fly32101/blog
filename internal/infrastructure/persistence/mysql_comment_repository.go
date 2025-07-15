package persistence

import (
	"database/sql"
	"fmt"

	"blog/internal/domain/entity"
	"blog/internal/domain/repository"
)

// MySQLCommentRepository MySQL评论存储库实现
type MySQLCommentRepository struct {
	db *sql.DB
}

// NewMySQLCommentRepository 创建MySQL评论存储库
func NewMySQLCommentRepository(conn *MySQLConnection) repository.CommentRepository {
	return &MySQLCommentRepository{
		db: conn.DB,
	}
}

// Create 创建评论
func (r *MySQLCommentRepository) Create(comment *entity.Comment) error {
	query := `INSERT INTO comments (post_id, content, author, created_at) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query, comment.PostID, comment.Content, comment.Author, comment.CreatedAt)
	if err != nil {
		return fmt.Errorf("创建评论失败: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("获取评论ID失败: %w", err)
	}

	comment.ID = uint(id)
	return nil
}

// GetByID 根据ID获取评论
func (r *MySQLCommentRepository) GetByID(id uint) (*entity.Comment, error) {
	query := `SELECT id, post_id, content, author, created_at FROM comments WHERE id = ?`
	row := r.db.QueryRow(query, id)

	comment := &entity.Comment{}
	err := row.Scan(&comment.ID, &comment.PostID, &comment.Content, &comment.Author, &comment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("评论不存在: %d", id)
		}
		return nil, fmt.Errorf("获取评论失败: %w", err)
	}

	return comment, nil
}

// GetByPostID 根据文章ID获取评论
func (r *MySQLCommentRepository) GetByPostID(postID uint) ([]*entity.Comment, error) {
	query := `SELECT id, post_id, content, author, created_at FROM comments WHERE post_id = ? ORDER BY created_at DESC`
	rows, err := r.db.Query(query, postID)
	if err != nil {
		return nil, fmt.Errorf("获取文章评论失败: %w", err)
	}
	defer rows.Close()

	var comments []*entity.Comment
	for rows.Next() {
		comment := &entity.Comment{}
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content, &comment.Author, &comment.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("读取评论数据失败: %w", err)
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历评论数据失败: %w", err)
	}

	return comments, nil
}

// Delete 删除评论
func (r *MySQLCommentRepository) Delete(id uint) error {
	query := `DELETE FROM comments WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("删除评论失败: %w", err)
	}
	return nil
}
