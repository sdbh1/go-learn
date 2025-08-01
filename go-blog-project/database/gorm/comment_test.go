package database_test

import (
	database "sdbh/database/gorm"
	"sdbh/database/model"
	"sdbh/test"
	"strconv"
	"testing"

	"gorm.io/gorm"
)

func init() {
	test.InitAll()
}

// 测试创建评论
func TestCreateComment(t *testing.T) {
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback()

	userID, blogID := createTestUserAndBlog(t, tx)
	commentContent := "测试评论内容"
	commentID, err := database.CreateComment(tx, uint(userID), blogID, commentContent)
	if err != nil {
		t.Fatalf("创建评论失败: %v", err)
	}

	commentID2, err2 := database.CreateComment(tx, uint(userID), blogID, commentContent)
	if err != nil && commentID2 == commentID {
		t.Fatalf("评论ID重复: %v", err2)
	}

	var comment model.Comment
	tx.First(&comment, "id = ?", commentID)
	if comment.Content != commentContent {
		t.Fatalf("评论内容不一致，期望: %s, 实际: %s", commentContent, comment.Content)
	}
}

// 测试获取评论列表
func TestGetComments(t *testing.T) {
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback()

	userID, blogID := createTestUserAndBlog(t, tx)
	for i := 0; i < 3; i++ {
		_, err := database.CreateComment(tx, uint(userID), blogID, "测试评论内容 "+strconv.Itoa(i))
		if err != nil {
			t.Fatalf("创建评论失败: %v", err)
		}
	}

	comments, err := database.GetCommentsByPostID(tx, blogID)
	if err != nil {
		t.Fatalf("获取评论列表失败: %v", err)
	}
	if len(comments) != 3 {
		t.Fatalf("获取的评论数量不一致，期望: 3, 实际: %d", len(comments))
	}
}

// 测试删除评论
func TestDeleteComment(t *testing.T) {
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback()

	userID, blogID := createTestUserAndBlog(t, tx)
	commentID, err := database.CreateComment(tx, uint(userID), blogID, "测试评论内容")
	if err != nil {
		t.Fatalf("创建评论失败: %v", err)
	}

	err = database.DeleteComment(tx, commentID, uint(userID))
	if err != nil {
		t.Fatalf("删除评论失败: %v", err)
	}

	err = database.DeleteComment(tx, commentID, uint(userID))
	if err == nil {
		t.Fatalf("重复删除评论了: %v", err)
	}

	var comment model.Comment
	result := tx.First(&comment, "id = ?", commentID)
	if result.Error != gorm.ErrRecordNotFound {
		t.Fatalf("评论未被正确删除")
	}
}

// 创建测试用的用户和博客
func createTestUserAndBlog(t *testing.T, tx *gorm.DB) (int, uint) {
	database.RegisterUser("testuser", "testpassword", tx)
	var user model.User = model.User{}
	result := tx.Model(&user).Where("name = ?", "testuser").First(&user)
	if result.Error != nil {
		t.Fatalf("用户未创建成功: %v", result.Error)
	}

	post := model.Post{
		UserId:  uint(user.Id),
		Title:   "测试标题",
		Content: "测试内容",
	}
	blogID, err := database.CreateBlog(tx, post)
	if err != nil {
		t.Fatalf("创建博客失败: %v", err)
	}
	return user.Id, blogID
}
