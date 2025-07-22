# 博客API接口文档

## 基础信息

- 基础URL: `http://localhost:8080/api`
- 所有POST请求的Content-Type均为: `application/json`
- 所有响应均为JSON格式

## 统一响应格式

所有API响应均使用以下统一格式：

```json
{
  "success": true/false,  // 请求是否成功
  "data": {},             // 成功时返回的数据
  "message": "",          // 成功或错误消息
  "error": ""             // 错误时返回的错误信息
}
```

## 文章接口

### 1. 创建文章

- **URL**: `/posts`
- **方法**: `POST`
- **请求体**:
```json
{
  "title": "文章标题",
  "content": "文章内容",
  "author": "作者名称",
  "title_url": "https://example.com/image.jpg",  // 可选，文章头图URL
  "category_ids": [1, 2]  // 可选，文章所属分类ID数组
}
```
- **成功响应** (201 Created):
```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "文章标题",
    "author": "作者名称",
    "title_url": "https://example.com/image.jpg",
    "view_count": 0,
    "created_at": "2023-07-15T13:45:30Z",
    "updated_at": "2023-07-15T13:45:30Z"
  },
  "message": "文章创建成功"
}
```

### 2. 获取所有文章

- **URL**: `/posts`
- **方法**: `GET`
- **成功响应** (200 OK):
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "title": "文章标题1",
      "author": "作者1",
      "title_url": "https://example.com/image1.jpg",
      "view_count": 15,
      "created_at": "2023-07-15T13:45:30Z",
      "updated_at": "2023-07-15T13:45:30Z"
    },
    {
      "id": 2,
      "title": "文章标题2",
      "author": "作者2",
      "title_url": "https://example.com/image2.jpg",
      "view_count": 8,
      "created_at": "2023-07-15T14:12:22Z",
      "updated_at": "2023-07-15T14:12:22Z"
    }
  ],
  "message": "获取文章列表成功"
}
```

### 3. 根据ID获取文章详情

- **URL**: `/posts/{id}`
- **方法**: `GET`
- **成功响应** (200 OK):
```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "文章标题",
    "content": "文章内容",
    "author": "作者名称",
    "title_url": "https://example.com/image.jpg",
    "view_count": 10,
    "created_at": "2023-07-15T13:45:30Z",
    "updated_at": "2023-07-15T13:45:30Z",
    "categories": [
      {
        "id": 1,
        "name": "技术",
        "description": "技术相关文章"
      }
    ]
  },
  "message": "获取文章详情成功"
}
```
- **失败响应** (404 Not Found):
```json
{
  "success": false,
  "data": null,
  "message": "文章不存在",
  "error": "文章不存在: 999"
}
```

### 4. 更新文章

- **URL**: `/posts/{id}`
- **方法**: `PUT`
- **请求体**:
```json
{
  "title": "更新后的标题",
  "content": "更新后的内容",
  "title_url": "https://example.com/new-image.jpg"  // 可选，更新后的头图URL
}
```
- **成功响应** (200 OK):
```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "更新后的标题",
    "author": "作者名称",
    "title_url": "https://example.com/new-image.jpg",
    "view_count": 5,
    "created_at": "2023-07-15T13:45:30Z",
    "updated_at": "2023-07-15T15:30:45Z"
  },
  "message": "文章更新成功"
}
```

### 5. 删除文章

- **URL**: `/posts/{id}`
- **方法**: `DELETE`
- **成功响应** (200 OK):
```json
{
  "success": true,
  "data": null,
  "message": "文章删除成功"
}
```

### 6. 根据分类获取文章

- **URL**: `/posts/category/{id}`
- **方法**: `GET`
- **成功响应** (200 OK):
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "title": "文章标题1",
      "author": "作者1",
      "title_url": "https://example.com/image1.jpg",
      "view_count": 15,
      "created_at": "2023-07-15T13:45:30Z",
      "updated_at": "2023-07-15T13:45:30Z"
    },
    {
      "id": 3,
      "title": "文章标题3",
      "author": "作者3",
      "title_url": "https://example.com/image3.jpg",
      "view_count": 3,
      "created_at": "2023-07-15T16:22:10Z",
      "updated_at": "2023-07-15T16:22:10Z"
    }
  ],
  "message": "获取分类文章成功"
}
```

## 评论接口

### 1. 创建评论

- **URL**: `/comments`
- **方法**: `POST`
- **请求体**:
```json
{
  "post_id": 1,
  "content": "评论内容",
  "author": "评论者名称"
}
```
- **成功响应** (201 Created):
```json
{
  "success": true,
  "data": {
    "id": 1,
    "post_id": 1,
    "content": "评论内容",
    "author": "评论者名称",
    "created_at": "2023-07-15T14:30:22Z"
  },
  "message": "评论创建成功"
}
```

### 2. 根据ID获取评论

- **URL**: `/comments/{id}`
- **方法**: `GET`
- **成功响应** (200 OK):
```json
{
  "success": true,
  "data": {
    "id": 1,
    "post_id": 1,
    "content": "评论内容",
    "author": "评论者名称",
    "created_at": "2023-07-15T14:30:22Z"
  },
  "message": "获取评论成功"
}
```

### 3. 获取文章的所有评论

- **URL**: `/comments/post/{id}`
- **方法**: `GET`
- **成功响应** (200 OK):
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "post_id": 1,
      "content": "第一条评论",
      "author": "评论者1",
      "created_at": "2023-07-15T14:30:22Z"
    },
    {
      "id": 2,
      "post_id": 1,
      "content": "第二条评论",
      "author": "评论者2",
      "created_at": "2023-07-15T15:12:05Z"
    }
  ],
  "message": "获取文章评论成功"
}
```

### 4. 删除评论

- **URL**: `/comments/{id}`
- **方法**: `DELETE`
- **成功响应** (200 OK):
```json
{
  "success": true,
  "data": null,
  "message": "评论删除成功"
}
```

## 分类接口

### 1. 创建分类

- **URL**: `/categories`
- **方法**: `POST`
- **请求体**:
```json
{
  "name": "分类名称",
  "description": "分类描述"
}
```
- **成功响应** (201 Created):
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "分类名称",
    "description": "分类描述"
  },
  "message": "分类创建成功"
}
```

### 2. 获取所有分类

- **URL**: `/categories`
- **方法**: `GET`
- **成功响应** (200 OK):
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "技术",
      "description": "技术相关文章"
    },
    {
      "id": 2,
      "name": "生活",
      "description": "生活相关文章"
    }
  ],
  "message": "获取分类列表成功"
}
```

### 3. 根据ID获取分类

- **URL**: `/categories/{id}`
- **方法**: `GET`
- **成功响应** (200 OK):
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "技术",
    "description": "技术相关文章"
  },
  "message": "获取分类详情成功"
}
```

### 4. 更新分类

- **URL**: `/categories/{id}`
- **方法**: `PUT`
- **请求体**:
```json
{
  "name": "更新后的分类名称",
  "description": "更新后的分类描述"
}
```
- **成功响应** (200 OK):
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "更新后的分类名称",
    "description": "更新后的分类描述"
  },
  "message": "分类更新成功"
}
```

### 5. 删除分类

- **URL**: `/categories/{id}`
- **方法**: `DELETE`
- **成功响应** (200 OK):
```json
{
  "success": true,
  "data": null,
  "message": "分类删除成功"
}
```

## 测试示例

以下是使用curl测试API的示例命令：

### 创建分类
```bash
curl -X POST http://localhost:8080/api/categories -H "Content-Type: application/json" -d '{"name":"技术","description":"技术相关文章"}'
```

### 创建文章
```bash
curl -X POST http://localhost:8080/api/posts -H "Content-Type: application/json" -d '{"title":"Go语言入门","content":"这是一篇Go语言入门文章","author":"张三","title_url":"https://example.com/go-tutorial.jpg","category_ids":[1]}'
```

### 获取所有文章
```bash
curl http://localhost:8080/api/posts
```

### 获取特定文章
```bash
curl http://localhost:8080/api/posts/1
```

### 更新文章
```bash
curl -X PUT http://localhost:8080/api/posts/1 -H "Content-Type: application/json" -d '{"title":"Go语言进阶","content":"这是一篇Go语言进阶文章","title_url":"https://example.com/go-advanced.jpg"}'
```

### 为文章添加评论
```bash
curl -X POST http://localhost:8080/api/comments -H "Content-Type: application/json" -d '{"post_id":1,"content":"很好的文章!","author":"李四"}'
```

### 获取文章的评论
```bash
curl http://localhost:8080/api/comments/post/1
```

## 错误处理

所有API在遇到错误时会返回相应的HTTP状态码和统一格式的错误信息：

- **400 Bad Request**: 请求参数错误
- **404 Not Found**: 请求的资源不存在
- **500 Internal Server Error**: 服务器内部错误

错误响应示例：
```json
{
  "success": false,
  "data": null,
  "message": "请求处理失败",
  "error": "错误信息描述"
}
``` 