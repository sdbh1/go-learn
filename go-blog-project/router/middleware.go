package router

import (
	handler "sdbh/handler/gin"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// registerGlobalMiddleware 注册全局中间件
func registerGlobalMiddleware(engine *gin.Engine) {
	// 全局 metric 中间件
	engine.Use(handler.Metric)

	// 监控路由
	engine.GET("/metrics", func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})
}

// GetAuthMiddleware 返回认证中间件
func GetAuthMiddleware() gin.HandlerFunc {
	return handler.Auth
}
