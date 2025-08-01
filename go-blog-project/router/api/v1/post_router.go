package v1

import (
	handler "sdbh/handler/gin"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(api *gin.RouterGroup) {
	// 应用认证中间件
	group := api.Group("/post")

	// 发布相关接口
	group.POST("/blog/list", handler.GetBlogList)

	group.POST("/blog/publish", handler.PublishBlog)

	group.POST("/blog/delete", handler.DeleteBlog)

	group.POST("/blog/:id/detail", handler.GetBlogDetail)

	group.POST("/blog/search/maxComment", handler.GetMaxCommentBlog)
}
