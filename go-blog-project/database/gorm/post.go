package database

import (
	"errors"
	"sdbh/database/model"
	"sdbh/logger"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// 修改 GetBlogList 函数，添加 db 参数
func GetBlogList(db *gorm.DB, page uint) ([]model.Post, error) {
	posts := []model.Post{}
	var singlePageNum uint = 3
	offset := (page - 1) * singlePageNum

	result := db.Debug().Offset(int(offset)).Limit(int(singlePageNum)).Find(&posts)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return posts, logger.Error("[database][GetBlogList] no Any Data ", "error", result.Error.Error())
	} else if result.Error != nil {
		return posts, logger.Error("[database][GetBlogList] fail ", "error", result.Error.Error())
	} else if posts == nil {
		return posts, logger.Error("[database][GetBlogList] fail ", "error", result.Error.Error())
	}

	logger.Info("[database][GetBlogList] success ", "pageTotalNum", len(posts))

	return posts, nil
}

// 修改 PublishBlog 函数，添加 db 参数
func PublishBlog(db *gorm.DB, uid, blogID uint, title string, content string) (uint, error) {
	post := model.Post{
		UserId:  uid,
		Title:   title,
		Content: content,
	}
	var err error
	if blogID == 0 {
		blogID, err = CreateBlog(db, post)
	} else {
		err = UpdateBlog(db, post, blogID)
	}

	if err != nil {
		return blogID, err
	}
	return blogID, nil
}

// 修改 CreateBlog 函数，添加 db 参数
func CreateBlog(db *gorm.DB, post model.Post) (uint, error) {
	tx := db.Debug().Create(&post)
	err := tx.Error
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			return post.ID, logger.Error("[database][CreateBlog] fail ", "error", mysqlErr.Error(), "errorNumber", mysqlErr.Number)
		}
		return post.ID, logger.Error("[database][CreateBlog] fail ", "error", err.Error())
	}
	return post.ID, logger.Info("[database][CreateBlog] success ", "id", post.ID)
}

// 修改 UpdateBlog 函数，添加 db 参数
func UpdateBlog(db *gorm.DB, newPost model.Post, blogID uint) error {
	var dbPosts model.Post

	result := db.Debug().Model(&dbPosts).First(&dbPosts, blogID)

	if result.Error != nil {
		return logger.Error("[database][UpdateBlog] fail ", "error", result.Error.Error())
	}

	if dbPosts.UserId != newPost.UserId {
		return logger.Error("[database][UpdateBlog] fail ", "error", "no author cant update")
	}

	dbPosts.Content = newPost.Content
	dbPosts.Title = newPost.Title
	result = db.Save(&dbPosts)

	if result.Error != nil {
		return logger.Error("[database][UpdateBlog] fail ", "error", result.Error.Error())
	}
	return nil
}

// 修改 DeleteBlog 函数，添加 db 参数
func DeleteBlog(db *gorm.DB, id, reqUid uint) error {
	var dbPosts model.Post

	result := db.Debug().Model(&dbPosts).First(&dbPosts, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return logger.Error("[database][DeleteBlog] fail ", "error", result.Error.Error())
	} else if result.Error != nil {
		return logger.Error("[database][DeleteBlog] fail ", "error", result.Error.Error())
	}

	if dbPosts.UserId != reqUid {
		return logger.Error("[database][DeleteBlog] fail ", "error", "no author cant delete")
	}
	result = db.Debug().Delete(&dbPosts)
	if result.Error != nil {
		return logger.Error("[database][DeleteBlog] fail ", "error", result.Error.Error())
	}
	return nil
}

// 修改 GetBlogDetail 函数，添加 db 参数
func GetBlogDetail(db *gorm.DB, id uint) (model.Post, error) {
	var post model.Post
	result := db.Debug().First(&post, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return post, logger.Error("[database][GetBlogByID] fail ", "error", "文章未找到")
	} else if result.Error != nil {
		return post, logger.Error("[database][GetBlogByID] fail ", "error", result.Error.Error())
	}
	return post, nil
}

// 修改 GetMaxCommentBlog 函数，添加 db 参数
func GetMaxCommentBlog(db *gorm.DB) (model.Post, error) {
	var post model.Post
	result := db.Model(&model.Post{}).
		Order("comment_num DESC").Limit(1).Find(&post)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return post, logger.Error("[database][GetMaxCommentBlog] fail ", "error", result.Error.Error())
	} else if result.Error != nil {
		return post, logger.Error("[database][GetMaxCommentBlog] fail ", "error", result.Error.Error())
	}

	return post, logger.Info("[database][GetMaxCommentBlog] success ", "id", post.ID)
}

func GetUserAllBriefPost(db *gorm.DB, uid uint) ([]*model.Post, error) {
	var posts []*model.Post
	tx := db.Model(&model.Post{}).
		Where("user_id = ?", uid).
		Select("id,user_id, title, create_time, comment_num, like_num").
		Find(&posts)

	if tx.Error != nil {
		return nil, logger.Error("[database][GetUserAllBriefPost] fail ", "error", tx.Error.Error())
	}

	return posts, logger.Info("[database][GetUserAllBriefPost] success ", "num", len(posts))
}
