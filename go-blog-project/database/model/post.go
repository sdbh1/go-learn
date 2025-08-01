package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID         uint           `gorm:"primarykey"`
	CreatedAt  time.Time      `gorm:"column:create_time" ` //创建时间
	UpdatedAt  time.Time      `gorm:"column:update_time" ` //更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:delete_time" ` //删除时间
	UserId     uint           `gorm:"column:user_id"`      //发布者id
	Title      string         `gorm:"column:title"`        //标题
	Content    string         `gorm:"column:content" `     //正文
	CommentNum int            `gorm:"column:comment_num"`
	LikeNum    int            `gorm:"column:like_num"`
}

// AfterCreate 在文章创建后更新用户的文章数量统计字段
func (p *Post) AfterCreate(tx *gorm.DB) error {
	var user User
	if err := tx.First(&user, p.UserId).Error; err != nil {
		return err
	}
	user.PostNum++
	return tx.Save(&user).Error
}
