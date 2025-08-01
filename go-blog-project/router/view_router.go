package router

import (
	"net/http"
	"sdbh/global"

	"github.com/gin-gonic/gin"
)

// RegisterViewRoutes 注册页面相关路由
func RegisterViewRoutes(engine *gin.Engine) {
	// 静态资源配置
	setupStaticResources(engine)

	// 页面路由
	engine.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})

	engine.GET("/regist", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "user_regist.html", nil)
	})

	// 首页重定向
	engine.GET("", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "news")
	})
}

// setupStaticResources 配置静态资源
func setupStaticResources(engine *gin.Engine) {
	// 静态文件目录
	engine.Static("/js", global.ProjectRootPath+"/views/js")
	engine.Static("/css", global.ProjectRootPath+"/views/css")
	engine.StaticFile("/favicon.ico", global.ProjectRootPath+"/views/img/dqq.png")

	// HTML模板
	engine.LoadHTMLGlob(global.ProjectRootPath + "/views/html/*")
}
