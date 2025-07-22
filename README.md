# 博客API项目

这是一个使用 Go 语言、Gin 框架和领域驱动设计（DDD）架构开发的博客 API 项目。

## 项目结构

```
blog/
├── cmd/                  # 应用程序入口
│   └── api/              # API 服务入口
├── internal/             # 内部包
│   ├── application/      # 应用层
│   ├── config/           # 配置
│   ├── domain/           # 领域层
│   │   ├── entity/       # 实体
│   │   ├── repository/   # 仓储接口
│   │   └── service/      # 领域服务
│   ├── infrastructure/   # 基础设施层
│   │   └── persistence/  # 持久化实现
│   └── interfaces/       # 接口层
│       ├── api/          # API 处理器
│       └── dto/          # 数据传输对象
├── pkg/                  # 公共包
│   ├── middleware/       # 中间件
│   └── utils/            # 工具函数
└── test/                 # 测试
    └── unit/             # 单元测试
```

## 数据库表结构

### posts 表（文章表）
```sql
CREATE TABLE IF NOT EXISTS posts (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author VARCHAR(100) NOT NULL,
    title_url VARCHAR(500),
    view_count INT UNSIGNED DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### comments 表（评论表）
```sql
CREATE TABLE IF NOT EXISTS comments (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    post_id INT UNSIGNED NOT NULL,
    content TEXT NOT NULL,
    author VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### categories 表（分类表）
```sql
CREATE TABLE IF NOT EXISTS categories (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### post_categories 表（文章分类关联表）
```sql
CREATE TABLE IF NOT EXISTS post_categories (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    post_id INT UNSIGNED NOT NULL,
    category_id INT UNSIGNED NOT NULL,
    UNIQUE KEY (post_id, category_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

## API 接口

### 文章接口
- `POST /api/posts` - 创建文章
- `GET /api/posts` - 获取所有文章
- `GET /api/posts/:id` - 根据ID获取文章
- `PUT /api/posts/:id` - 更新文章
- `DELETE /api/posts/:id` - 删除文章
- `GET /api/posts/category/:id` - 根据分类获取文章

### 评论接口
- `POST /api/comments` - 创建评论
- `GET /api/comments/:id` - 根据ID获取评论
- `GET /api/comments/post/:id` - 获取文章的所有评论
- `DELETE /api/comments/:id` - 删除评论

### 分类接口
- `POST /api/categories` - 创建分类
- `GET /api/categories` - 获取所有分类
- `GET /api/categories/:id` - 根据ID获取分类
- `PUT /api/categories/:id` - 更新分类
- `DELETE /api/categories/:id` - 删除分类

## 如何运行

1. 确保已安装 Go 环境（推荐 Go 1.21 或更高版本）
2. 克隆项目到本地
3. 配置数据库信息（修改 `internal/config/config.go` 文件）
4. 运行项目：
   ```bash
   go run cmd/api/main.go
   ```
5. 访问 API：http://localhost:8080/api/

## 单元测试

运行单元测试：
```bash
go test ./test/unit/...
``` 