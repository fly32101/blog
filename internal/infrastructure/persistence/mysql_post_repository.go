package persistence

import (
	"database/sql"
	"fmt"
	"time"

	"blog/internal/domain/entity"
	"blog/internal/domain/repository"
)

// MySQLPostRepository MySQL文章存储库实现
type MySQLPostRepository struct {
	db *sql.DB
}

// NewMySQLPostRepository 创建MySQL文章存储库
func NewMySQLPostRepository(conn *MySQLConnection) repository.PostRepository {
	return &MySQLPostRepository{
		db: conn.DB,
	}
}

// Create 创建文章
func (r *MySQLPostRepository) Create(post *entity.Post) error {
	query := `INSERT INTO posts (title, content, author, title_url, view_count, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, post.Title, post.Content, post.Author, post.TitleURL, post.ViewCount, post.CreatedAt, post.UpdatedAt)
	if err != nil {
		return fmt.Errorf("创建文章失败: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("获取文章ID失败: %w", err)
	}

	post.ID = uint(id)
	return nil
}

// GetByID 根据ID获取文章
func (r *MySQLPostRepository) GetByID(id uint) (*entity.Post, error) {
	query := `SELECT id, title, content, author, title_url, view_count, created_at, updated_at FROM posts WHERE id = ?`
	row := r.db.QueryRow(query, id)

	post := &entity.Post{}
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.TitleURL, &post.ViewCount, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("文章不存在: %d", id)
		}
		return nil, fmt.Errorf("获取文章失败: %w", err)
	}

	return post, nil
}

// GetAll 获取所有文章
func (r *MySQLPostRepository) GetAll() ([]*entity.Post, error) {
	query := `SELECT id, title, content, author, title_url, view_count, created_at, updated_at FROM posts ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("获取所有文章失败: %w", err)
	}
	defer rows.Close()

	var posts []*entity.Post
	for rows.Next() {
		post := &entity.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.TitleURL, &post.ViewCount, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("读取文章数据失败: %w", err)
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历文章数据失败: %w", err)
	}

	return posts, nil
}

// Update 更新文章
func (r *MySQLPostRepository) Update(post *entity.Post) error {
	query := `UPDATE posts SET title = ?, content = ?, title_url = ?, view_count = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.Exec(query, post.Title, post.Content, post.TitleURL, post.ViewCount, time.Now(), post.ID)
	if err != nil {
		return fmt.Errorf("更新文章失败: %w", err)
	}
	return nil
}

// Delete 删除文章
func (r *MySQLPostRepository) Delete(id uint) error {
	query := `DELETE FROM posts WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("删除文章失败: %w", err)
	}
	return nil
}

// GetByCategory 根据分类获取文章
func (r *MySQLPostRepository) GetByCategory(categoryID uint) ([]*entity.Post, error) {
	query := `
		SELECT p.id, p.title, p.content, p.author, p.title_url, p.view_count, p.created_at, p.updated_at
		FROM posts p
		JOIN post_categories pc ON p.id = pc.post_id
		WHERE pc.category_id = ?
		ORDER BY p.created_at DESC
	`
	rows, err := r.db.Query(query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("获取分类文章失败: %w", err)
	}
	defer rows.Close()

	var posts []*entity.Post
	for rows.Next() {
		post := &entity.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.TitleURL, &post.ViewCount, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("读取分类文章数据失败: %w", err)
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历分类文章数据失败: %w", err)
	}

	return posts, nil
}
