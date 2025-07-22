package api

import (
	"github.com/gin-gonic/gin"
)

// Router API路由器
type Router struct {
	engine          *gin.Engine
	postHandler     *PostHandler
	commentHandler  *CommentHandler
	categoryHandler *CategoryHandler
}

// NewRouter 创建路由器
func NewRouter(
	engine *gin.Engine,
	postHandler *PostHandler,
	commentHandler *CommentHandler,
	categoryHandler *CategoryHandler,
) *Router {
	return &Router{
		engine:          engine,
		postHandler:     postHandler,
		commentHandler:  commentHandler,
		categoryHandler: categoryHandler,
	}
}

// SetupRoutes 设置路由
func (r *Router) SetupRoutes() {
	api := r.engine.Group("/api")

	// 注册各个处理器的路由
	r.postHandler.Register(api)
	r.commentHandler.Register(api)
	r.categoryHandler.Register(api)
}

// Run 运行服务器
func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}
