package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:create_time" ` //创建时间
	UpdatedAt time.Time      `gorm:"column:update_time" ` //更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:delete_time" ` //删除时间
	UserId    uint           //发布者id
	User      User           `gorm:"-" ` //数据库里没有这一列
	PostID    uint           //标题
	Content   string         `gorm:"column:content" ` //正文
}

// AfterCreate 在评论创建后检查文章的评论数量并更新评论状态
func (c *Comment) AfterCreate(tx *gorm.DB) error {
	var post Post
	if err := tx.First(&post, c.PostID).Error; err != nil {
		return err
	}
	var commentCount int64 = 0

	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return err
	}

	post.CommentNum = int(commentCount)
	if err := tx.Save(&post).Error; err != nil {
		return err
	}
	return nil
}

// AfterDelete 在评论删除后检查文章的评论数量并更新评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var post Post
	if err := tx.First(&post, c.PostID).Error; err != nil {
		return err
	}
	var commentCount int64 = 0

	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return err
	}

	post.CommentNum = int(commentCount)
	if err := tx.Save(&post).Error; err != nil {
		return err
	}
	return nil
}
