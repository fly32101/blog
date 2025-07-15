package main

import (
	"database/sql"
	"fmt"
	"log"

	"blog/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 使用与config.go相同的配置
	cfg := config.NewConfig()
	dsn := cfg.Database.GetDSN()

	// 连接到MySQL
	db, err := sql.Open(cfg.Database.Driver, dsn)
	if err != nil {
		log.Fatalf("无法打开数据库连接: %v", err)
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	fmt.Println("数据库连接成功!")
	fmt.Printf("使用DSN: %s\n", dsn)

	// 创建文章表
	_, err = db.Exec(`
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
		log.Fatalf("创建文章表失败: %v", err)
	}
	fmt.Println("文章表创建成功!")

	// 创建评论表
	_, err = db.Exec(`
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
		log.Fatalf("创建评论表失败: %v", err)
	}
	fmt.Println("评论表创建成功!")

	// 创建分类表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL UNIQUE,
			description VARCHAR(255)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`)
	if err != nil {
		log.Fatalf("创建分类表失败: %v", err)
	}
	fmt.Println("分类表创建成功!")

	// 创建文章分类关联表
	_, err = db.Exec(`
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
		log.Fatalf("创建文章分类关联表失败: %v", err)
	}
	fmt.Println("文章分类关联表创建成功!")
}
