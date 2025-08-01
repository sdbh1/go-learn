package handler

import (
	"net/http"
	database "sdbh/database/gorm"
	"sdbh/global"
	"sdbh/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCommentHandler 创建评论接口
func CreateCommentHandler(ctx *gin.Context) {
	loginUID, err := GetUIDByCtx(ctx)
	if err != nil {
		logger.Error("[Handler][CreateComment] user not logged in: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "用户未登录"})
		return
	}

	var input struct {
		PostID  uint   `json:"postID"`
		Content string `json:"content"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentID, err := database.CreateComment(global.BlogDB, loginUID, input.PostID, input.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "评论创建成功", "id": commentID})
}

// GetCommentsHandler 获取文章评论列表接口
func GetCommentsHandler(ctx *gin.Context) {

	blogIDStr := ctx.Param("id")

	// 检查 blogID 是否为正整数
	blogIDInt, err := strconv.Atoi(blogIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "blogID must be a positive integer"})
		return
	}

	if blogIDInt <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "blogID must be a positive integer"})
		return
	}

	comments, err := database.GetCommentsByPostID(global.BlogDB, uint(blogIDInt))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "获取评论成功", "data": comments})
}

// DeleteCommentHandler 删除评论接口
func DeleteCommentHandler(ctx *gin.Context) {
	loginUID, err := GetUIDByCtx(ctx)
	if err != nil {
		logger.Error("[Handler][DeleteComment] user not logged in: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "用户未登录"})
		return
	}

	var input struct {
		CommentID uint `json:"commentID"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.DeleteComment(global.BlogDB, input.CommentID, loginUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "评论删除成功"})
}
