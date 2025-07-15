package persistence

import (
	"database/sql"
	"fmt"
	"log"

	"blog/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLConnection 数据库连接
type MySQLConnection struct {
	DB *sql.DB
}

// NewMySQLConnection 创建新的数据库连接
func NewMySQLConnection(config *config.DatabaseConfig) (*MySQLConnection, error) {
	db, err := sql.Open(config.Driver, config.GetDSN())
	if err != nil {
		return nil, err
	}

	// 测试连接
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("数据库连接成功")
	return &MySQLConnection{DB: db}, nil
}

// InitTables 初始化数据库表
func (conn *MySQLConnection) InitTables() error {
	// 创建文章表
	_, err := conn.DB.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			content TEXT NOT NULL,
			author VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`)
	if err != nil {
		return fmt.Errorf("创建文章表失败: %w", err)
	}

	// 创建评论表
	_, err = conn.DB.Exec(`
		CREATE TABLE IF NOT EXISTS comments (
			id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			post_id INT UNSIGNED NOT NULL,
			content TEXT NOT NULL,
			author VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`)
	if err != nil {
		return fmt.Errorf("创建评论表失败: %w", err)
	}

	// 创建分类表
	_, err = conn.DB.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL UNIQUE,
			description VARCHAR(255)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`)
	if err != nil {
		return fmt.Errorf("创建分类表失败: %w", err)
	}

	// 创建文章分类关联表
	_, err = conn.DB.Exec(`
		CREATE TABLE IF NOT EXISTS post_categories (
			id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			post_id INT UNSIGNED NOT NULL,
			category_id INT UNSIGNED NOT NULL,
			UNIQUE KEY (post_id, category_id),
			FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
			FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`)
	if err != nil {
		return fmt.Errorf("创建文章分类关联表失败: %w", err)
	}

	log.Println("数据库表初始化成功")
	return nil
}

// Close 关闭数据库连接
func (conn *MySQLConnection) Close() error {
	if conn.DB != nil {
		return conn.DB.Close()
	}
	return nil
}
