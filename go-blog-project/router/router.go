package router

import (
	"time"

	v1 "sdbh/router/api/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

// InitRouter 初始化路由引擎并启动服务
func InitRouter() *gin.Engine {
	// 初始化Gin引擎
	engine = gin.Default()

	// 配置全局中间件
	setupGlobalMiddleware()

	// 注册所有路由
	registerRoutes()

	return engine
}

// setupGlobalMiddleware 配置全局中间件
func setupGlobalMiddleware() {
	// CORS配置
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5678"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 注册全局中间件
	registerGlobalMiddleware(engine)
}

// registerRoutes 注册所有路由
func registerRoutes() {
	// 注册页面路由
	RegisterViewRoutes(engine)

	// 注册认证路由
	RegisterAuthRoutes(engine)

	// 注册API路由（按版本）
	v1.RegisterRoutes(engine.Group("/api/v1"))
}
