package v1

import (
	handler "sdbh/handler/gin"

	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(api *gin.RouterGroup) {
	// 应用认证中间件
	group := api.Group("/comment")

	// 评论相关接口
	group.POST("/blog/:id/list", handler.GetCommentsHandler)

	group.POST("/blog/publish", handler.CreateCommentHandler)

	group.POST("/blog/delete", handler.DeleteCommentHandler)

}
