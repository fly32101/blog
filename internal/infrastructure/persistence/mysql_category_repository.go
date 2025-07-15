package persistence

import (
	"database/sql"
	"fmt"

	"blog/internal/domain/entity"
	"blog/internal/domain/repository"
)

// MySQLCategoryRepository MySQL分类存储库实现
type MySQLCategoryRepository struct {
	db *sql.DB
}

// NewMySQLCategoryRepository 创建MySQL分类存储库
func NewMySQLCategoryRepository(conn *MySQLConnection) repository.CategoryRepository {
	return &MySQLCategoryRepository{
		db: conn.DB,
	}
}

// Create 创建分类
func (r *MySQLCategoryRepository) Create(category *entity.Category) error {
	query := `INSERT INTO categories (name, description) VALUES (?, ?)`
	result, err := r.db.Exec(query, category.Name, category.Description)
	if err != nil {
		return fmt.Errorf("创建分类失败: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("获取分类ID失败: %w", err)
	}

	category.ID = uint(id)
	return nil
}

// GetByID 根据ID获取分类
func (r *MySQLCategoryRepository) GetByID(id uint) (*entity.Category, error) {
	query := `SELECT id, name, description FROM categories WHERE id = ?`
	row := r.db.QueryRow(query, id)

	category := &entity.Category{}
	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("分类不存在: %d", id)
		}
		return nil, fmt.Errorf("获取分类失败: %w", err)
	}

	return category, nil
}

// GetAll 获取所有分类
func (r *MySQLCategoryRepository) GetAll() ([]*entity.Category, error) {
	query := `SELECT id, name, description FROM categories ORDER BY name`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("获取所有分类失败: %w", err)
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		category := &entity.Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, fmt.Errorf("读取分类数据失败: %w", err)
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历分类数据失败: %w", err)
	}

	return categories, nil
}

// Update 更新分类
func (r *MySQLCategoryRepository) Update(category *entity.Category) error {
	query := `UPDATE categories SET name = ?, description = ? WHERE id = ?`
	_, err := r.db.Exec(query, category.Name, category.Description, category.ID)
	if err != nil {
		return fmt.Errorf("更新分类失败: %w", err)
	}
	return nil
}

// Delete 删除分类
func (r *MySQLCategoryRepository) Delete(id uint) error {
	query := `DELETE FROM categories WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("删除分类失败: %w", err)
	}
	return nil
}

// AddPostToCategory 将文章添加到分类
func (r *MySQLCategoryRepository) AddPostToCategory(postID, categoryID uint) error {
	query := `INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)`
	_, err := r.db.Exec(query, postID, categoryID)
	if err != nil {
		return fmt.Errorf("添加文章到分类失败: %w", err)
	}
	return nil
}

// RemovePostFromCategory 从分类中移除文章
func (r *MySQLCategoryRepository) RemovePostFromCategory(postID, categoryID uint) error {
	query := `DELETE FROM post_categories WHERE post_id = ? AND category_id = ?`
	_, err := r.db.Exec(query, postID, categoryID)
	if err != nil {
		return fmt.Errorf("从分类中移除文章失败: %w", err)
	}
	return nil
}

// GetCategoriesByPostID 获取文章的所有分类
func (r *MySQLCategoryRepository) GetCategoriesByPostID(postID uint) ([]*entity.Category, error) {
	query := `
		SELECT c.id, c.name, c.description 
		FROM categories c
		JOIN post_categories pc ON c.id = pc.category_id
		WHERE pc.post_id = ?
		ORDER BY c.name
	`
	rows, err := r.db.Query(query, postID)
	if err != nil {
		return nil, fmt.Errorf("获取文章分类失败: %w", err)
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		category := &entity.Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, fmt.Errorf("读取文章分类数据失败: %w", err)
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历文章分类数据失败: %w", err)
	}

	return categories, nil
}
