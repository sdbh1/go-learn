package database_test

import (
	database "sdbh/database/gorm"
	"sdbh/database/model"
	"sdbh/test"
	"testing"

	"gorm.io/gorm"
)

func init() {
	test.InitAll()
}

// 测试发布博客
func TestPublishBlog(t *testing.T) {
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback()

	blogID, userID := CreateBlog(t, tx, "testuser", "testtitle", "testcontent")

	// 这里可添加查询逻辑验证博客是否创建成功
	var post model.Post
	tx.First(&post, "id = ?", blogID)

	//测试再次创建博客
	blogID, err := database.PublishBlog(tx, uint(userID), blogID, "测试标题2", "测试内容2")
	if err != nil {
		t.Fatalf("发布博客失败: %v", err)
	}

	tx.First(&post, "id = ?", blogID)

	//比对博客内容是否和第二次修改后的一样
	tx.First(&post, "title = ?", "测试标题2")
	if post.Content != "测试内容2" {
		t.Fatalf("博客内容不一致")
	}
}

// 测试获取博客列表
func TestGetBlogList(t *testing.T) {
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback()

	CreateBlog(t, tx, "testuser", "testtitle", "testcontent")
	CreateBlog(t, tx, "testuser1", "testtitle1", "testcontent1")
	CreateBlog(t, tx, "testuser2", "testtitle2", "testcontent2")
	CreateBlog(t, tx, "testuser", "testtitle3", "testcontent3")

	// 测试正常获取博客列表
	posts, err := database.GetBlogList(tx, 1)
	if err != nil {
		t.Fatalf("获取博客列表失败: %v", err)
	}
	t.Logf("获取到 %d 篇博客", len(posts))
}

// 测试删除博客
func TestDeleteBlog(t *testing.T) {
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback()

	database.RegisterUser("testuser", "testpassword", tx)

	blogID, userID := CreateBlog(t, tx, "testuser", "testtitle", "testcontent")

	err := database.DeleteBlog(tx, blogID, uint(userID)+1)
	if err == nil {
		t.Fatalf("博客被非发布者删除了: %v", err)
	}

	err = database.DeleteBlog(tx, blogID, uint(userID))
	if err != nil {
		t.Fatalf("删除博客失败: %v", err)
	}
}

// 测试获取博客详情
func TestGetBlogDetail(t *testing.T) {
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback()

	blogID, _ := CreateBlog(t, tx, "testuser", "详情测试标题", "testcontent")

	detail, err := database.GetBlogDetail(tx, blogID)
	if err != nil {
		t.Fatalf("获取博客详情失败: %v", err)
	}
	if detail.Title != "详情测试标题" {
		t.Errorf("获取的博客标题不匹配，期望: %s, 实际: %s", "详情测试标题", detail.Title)
	}
}

// 测试获取评论最多的博客
func TestGetMaxCommentBlog(t *testing.T) {
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback()

	post, err := database.GetMaxCommentBlog(tx)
	if err != nil {
		t.Fatalf("获取评论最多的博客失败: %v", err)
	}
	t.Logf("获取到评论最多的博客 ID: %d", post.ID)
}

func TestGetUserAllBriefPost(t *testing.T) {
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback()

	blogID, _ := CreateBlog(t, tx, "testuser", "testtitle", "testcontent")
	blogID1, _ := CreateBlog(t, tx, "testuser", "testtitle1", "testcontent1")
	blogID2, _ := CreateBlog(t, tx, "testuser", "testtitle2", "testcontent2")
	blogID3, _ := CreateBlog(t, tx, "testuser", "testtitle3", "testcontent3")

	user := database.GetUserByName("testuser", tx)
	if user == nil {
		t.Fatalf("获取用户失败")
	}

	database.CreateComment(tx, uint(user.Id), blogID, "testcomment1")

	database.CreateComment(tx, uint(user.Id), blogID1, "testcomment2")

	database.CreateComment(tx, uint(user.Id), blogID2, "testcomment3")

	database.CreateComment(tx, uint(user.Id), blogID3, "testcomment4")

	database.CreateComment(tx, uint(user.Id), blogID1, "testcomment5")

	posts, err := database.GetUserAllBriefPost(tx, uint(user.Id))

	if err != nil {
		t.Fatalf("获取用户博客失败: %v", err)
	}

	if len(posts) != 4 {
		t.Fatalf("获取的博客数量不一致，期望: 4, 实际: %d", len(posts))
	}
	t.Logf("Number of posts retrieved: %d\n", len(posts))

	for i := 0; i < len(posts); i++ {
		post := posts[i]
		t.Logf("UID = %d, title = %s, commentNum = %d, likeNum = %d\n", post.UserId, post.Title, post.CommentNum, post.LikeNum)
	}
}

func CreateBlog(t *testing.T, tx *gorm.DB, userName, title, content string) (uint, int) {

	user := database.GetUserByName(userName, tx)

	if user == nil {
		database.RegisterUser(userName, "testpassword", tx)

		result := tx.Model(&user).Where("name = ?", userName).First(&user)
		if result.Error != nil {
			t.Fatalf("用户未创建成功: %v", result.Error)
		}
	}

	post := model.Post{
		UserId:  uint(user.Id),
		Title:   title,
		Content: content,
	}
	blogID, err := database.CreateBlog(tx, post)
	if err != nil {
		t.Fatalf("创建博客失败: %v", err)
	}
	return blogID, user.Id
}
