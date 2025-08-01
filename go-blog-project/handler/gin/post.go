package handler

import (
	"net/http" // 修改导入路径
	database "sdbh/database/gorm"
	"sdbh/global"
	"sdbh/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PublishBlog(ctx *gin.Context) {

	loginUID, err := GetUIDByCtx(ctx)

	if err != nil {
		logger.Error("[DataBase][Out] user not already login :", loginUID)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		BlogID  uint   `json:"blogID"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogID, err := database.PublishBlog(global.BlogDB, loginUID, input.BlogID, input.Title, input.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "发布成功", "id": blogID})
}

func GetBlogList(ctx *gin.Context) {
	var input struct {
		Page uint `json:"page"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list, err := database.GetBlogList(global.BlogDB, input.Page) // 修改变量名

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "获取成功", "data": list})
}

func DeleteBlog(ctx *gin.Context) {

	loginUID, err := GetUIDByCtx(ctx)

	if err != nil {
		logger.Error("[DataBase][Out] user not already login :", loginUID)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	var input struct {
		ID uint `json:"id"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.DeleteBlog(global.BlogDB, input.ID, loginUID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "delete success!"})
}

func GetBlogDetail(ctx *gin.Context) {

	blogID := ctx.Param("id")

	blogIDInt, err := strconv.Atoi(blogID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	if blogIDInt <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "blogID must be greater than 0"})
		return
	}

	post, err := database.GetBlogDetail(global.BlogDB, uint(blogIDInt))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "get success!", "data": post})
}

func GetMaxCommentBlog(ctx *gin.Context) {

	post, err := database.GetMaxCommentBlog(global.BlogDB)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "get success!", "data": post})
}
