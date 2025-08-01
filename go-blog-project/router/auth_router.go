package router

import (
	handler "sdbh/handler/gin"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes 注册认证相关路由
func RegisterAuthRoutes(engine *gin.Engine) {
	authGroup := engine.Group("/auth")

	// 登录相关
	authGroup.POST("/login/submit", handler.Login)

	// 注册相关
	authGroup.POST("/register/submit", handler.RegisterUser)

	// 注销相关
	authGroup.POST("/login/logout", handler.LogOut)
}
