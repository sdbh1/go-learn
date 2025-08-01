package v1

import (
	handler "sdbh/handler/gin"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有v1版本的API路由
func RegisterRoutes(group *gin.RouterGroup) {
	group.Use(handler.Auth)
	// 注册用户相关路由
	RegisterUserRoutes(group)
	// 注册博客相关路由
	RegisterPostRoutes(group)
	// 注册评论相关路由
	RegisterCommentRoutes(group)
}
