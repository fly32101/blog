package main

import (
	"fmt"
	"log"

	"blog/internal/application"
	"blog/internal/config"
	"blog/internal/domain/service"
	"blog/internal/infrastructure/persistence"
	"blog/internal/interfaces/api"
	"blog/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg := config.NewConfig()

	// 初始化数据库连接
	dbConn, err := persistence.NewMySQLConnection(&cfg.Database)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer dbConn.Close()

	// 初始化数据库表
	err = dbConn.InitTables()
	if err != nil {
		log.Fatalf("初始化数据库表失败: %v", err)
	}

	// 初始化存储库
	postRepo := persistence.NewMySQLPostRepository(dbConn)
	commentRepo := persistence.NewMySQLCommentRepository(dbConn)
	categoryRepo := persistence.NewMySQLCategoryRepository(dbConn)

	// 初始化领域服务
	postService := service.NewPostService(postRepo, categoryRepo)
	commentService := service.NewCommentService(commentRepo, postRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	// 初始化应用服务
	postApp := application.NewPostApp(postService, categoryService)
	commentApp := application.NewCommentApp(commentService)
	categoryApp := application.NewCategoryApp(categoryService)

	// 初始化处理器
	postHandler := api.NewPostHandler(postApp)
	commentHandler := api.NewCommentHandler(commentApp)
	categoryHandler := api.NewCategoryHandler(categoryApp)

	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode)

	// 创建Gin引擎
	engine := gin.New()

	// 添加中间件
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())
	engine.Use(middleware.CORS())

	// 注册路由
	router := api.NewRouter(postHandler, commentHandler, categoryHandler)
	router.SetupRoutes()

	// 启动服务器
	port := cfg.Server.Port
	log.Printf("服务器启动在 http://localhost:%s", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
