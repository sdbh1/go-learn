package database

import (
	"errors"
	"sdbh/database/model"
	"sdbh/logger"

	"gorm.io/gorm"
)

// CreateComment 创建评论
func CreateComment(db *gorm.DB, uid, postID uint, content string) (uint, error) {
	comment := model.Comment{
		UserId:  uid,
		PostID:  postID,
		Content: content,
	}

	tx := db.Create(&comment)
	err := tx.Error

	if err != nil {
		return 0, logger.Error("[database][CreateComment] fail ", "error", err.Error())
	}
	logger.Info("[database][CreateComment] success ", "id", comment.ID)
	return comment.ID, nil
}

// GetCommentsByPostID 获取某篇文章的所有评论
func GetCommentsByPostID(db *gorm.DB, postID uint) ([]model.Comment, error) {
	comments := []model.Comment{}
	result := db.Debug().Where("post_id = ?", postID).Find(&comments)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return comments, logger.Error("[database][GetCommentsByPostID] no Any Data ", "error", result.Error.Error())
	} else if result.Error != nil {
		return comments, logger.Error("[database][GetCommentsByPostID] fail ", "error", result.Error.Error())
	}

	logger.Info("[database][GetCommentsByPostID] success ", "count", len(comments))
	return comments, nil
}

// DeleteComment 删除评论
func DeleteComment(db *gorm.DB, commentID, reqUid uint) error {
	var comment model.Comment
	result := db.Debug().First(&comment, commentID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return logger.Error("[database][DeleteComment] fail ", "error", "评论未找到")
	} else if result.Error != nil {
		return logger.Error("[database][DeleteComment] fail ", "error", result.Error.Error())
	}

	if comment.UserId != reqUid {
		return logger.Error("[database][DeleteComment] fail ", "error", "无权限删除该评论")
	}

	result = db.Debug().Delete(&comment)
	if result.Error != nil {
		return logger.Error("[database][DeleteComment] fail ", "error", result.Error.Error())
	}

	logger.Info("[database][DeleteComment] success ", "id", commentID)
	return nil
}
