package database_test

import (
	"crypto/md5"
	"encoding/hex"
	database "sdbh/database/gorm"
	"sdbh/test"
	"testing"
)

var (
	uid uint = 1
)

func hash(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	digest := hasher.Sum(nil)
	return hex.EncodeToString(digest) //md5的输出是128bit，十六进制编码之后长度是32
}

func init() {
	test.InitAll()
}

func TestRegisterUser(t *testing.T) {
	// 开启事务
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback() // 测试结束后回滚事务

	// 执行注册测试
	err := database.RegisterUser("testuser", "passwordhash", tx)
	if err != nil {
		t.Fatalf("首次注册失败: %v", err)
	}

	// 再次注册应该失败
	err = database.RegisterUser("testuser", "passwordhash", tx)
	if err == nil {
		t.Fatal("重复注册应该失败，但却成功了")
	}
}

func TestGetUserById(t *testing.T) {
	// 准备测试数据
	username := "testuser"
	password := hash("123456")

	// 开启事务
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback() // 测试结束后回滚事务

	// 注册测试用户
	err := database.RegisterUser(username, password, tx)
	if err != nil {
		t.Fatalf("注册测试用户失败: %v", err)
	}

	// 获取用户ID
	user := database.GetUserByName(username, tx)
	if user == nil {
		t.Fatalf("获取测试用户失败")
	}
	uid := uint(user.Id)

	// 测试正常获取用户
	foundUser := database.GetUserById(uid, tx)
	if foundUser == nil {
		t.Fatalf("could not get user by id %d", uid)
	}

	// 验证获取到的用户信息正确
	if foundUser.Id != int(uid) {
		t.Fatalf("获取到错误的用户，期望ID: %d, 实际ID: %d", uid, foundUser.Id)
	}
	if foundUser.UserName != username {
		t.Fatalf("获取到错误的用户名，期望: %s, 实际: %s", username, foundUser.UserName)
	}

	// 测试获取不存在的用户
	var tmpUid uint = 999999 // 一个不存在的用户ID
	user = database.GetUserById(tmpUid, tx)
	if user != nil {
		t.Fatalf("获取到不存在的用户，用户ID: %d, 用户信息: %v", tmpUid, *user)
	}
}

func TestGetUserByName(t *testing.T) {

	// 准备测试数据
	username := "testuser"
	password := hash("123456")

	// 开启事务
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback() // 测试结束后回滚事务

	// 注册测试用户
	err := database.RegisterUser(username, password, tx)
	if err != nil {
		t.Fatalf("注册测试用户失败: %v", err)
	}

	user := database.GetUserByName(username, tx)
	if user == nil {
		t.Fail()
	}

	user = database.GetUserByName("ok", tx)
	if user != nil {
		t.Fail()
	}
}

func TestUpdateUserName(t *testing.T) {
	// 准备测试数据
	username := "testuser"
	oldDisplayName := "默认名字"
	newDisplayName := "修改后名字"
	password := hash("123456")

	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback() // 测试结束后回滚事务
	// 注册测试用户
	err := database.RegisterUser(username, password, tx)
	if err != nil {
		t.Fatalf("注册测试用户失败: %v", err)
	}

	// 获取用户ID
	user := database.GetUserByName(username, tx)
	if user == nil {
		t.Fatalf("获取测试用户失败")
	}
	uid := uint(user.Id)

	// 测试更新用户名
	err = database.UpdateUserSingleField(uid, "display_name", oldDisplayName, newDisplayName, tx, true)
	if err != nil {
		t.Fatalf("更新用户名失败: %v", err)
	}

	// 验证用户名是否更新成功
	updatedUser := database.GetUserById(uid, tx)
	if updatedUser == nil {
		t.Fatalf("获取更新后的用户失败")
	}
	if updatedUser.DisplayName != newDisplayName {
		t.Fatalf("用户名更新失败: 期望 %s, 实际 %s", newDisplayName, updatedUser.UserName)
	}

	// 测试更新为相同的用户名
	err = database.UpdateUserSingleField(uid, "name", newDisplayName, newDisplayName, tx, true)
	if err == nil {
		t.Fatalf("更新为相同的用户名应该失败")
	}

	// 测试使用错误的旧值更新
	err = database.UpdateUserSingleField(uid, "name", "wrongoldname", "newname2", tx, true)
	if err == nil {
		t.Fatalf("使用错误的旧值更新应该失败")
	}
}

func TestUpdatePassword(t *testing.T) {
	// 准备测试数据
	username := "testuser2"
	password := hash("123456")
	newPassword := hash("654321")
	// 开启事务
	tx := database.Db.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback() // 测试结束后回滚事务
	// 注册测试用户
	err := database.RegisterUser(username, password, tx)
	if err != nil {
		t.Fatalf("注册测试用户失败: %v", err)
	}

	// 获取用户ID
	user := database.GetUserByName(username, tx)
	if user == nil {
		t.Fatalf("获取测试用户失败")
	}
	uid := uint(user.Id)

	// 测试更新密码
	err = database.UpdateUserSingleField(uid, "password", password, newPassword, tx, true)
	if err != nil {
		t.Fatalf("更新密码失败: %v", err)
	}

	// 验证密码是否更新成功
	updatedUser := database.GetUserById(uid, tx)
	if updatedUser == nil {
		t.Fatalf("获取更新后的用户失败")
	}
	if updatedUser.Password != newPassword {
		t.Fatalf("密码更新失败: 期望 %s, 实际 %s", newPassword, updatedUser.Password)
	}
}
