package v1

import (
	handler "sdbh/handler/gin"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(api *gin.RouterGroup) {
	// 应用认证中间件
	group := api.Group("/user")

	// 用户相关接口
	group.PUT("/update/Password", handler.UpdatePassword)

	group.PUT("/update/DisplayName", handler.UpdateDisplayName)
}
