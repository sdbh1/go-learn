package database

import (
	"errors"
	"log/slog"
	"sdbh/database/model"
	"sdbh/logger"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func GetUserByName(userName string, db *gorm.DB) *model.User {
	var user model.User
	tx := db.Where("name = ?", userName).First(&user)

	if mysqlErr, ok := tx.Error.(*mysql.MySQLError); ok {
		slog.Error("user not find :", "username", mysqlErr.Number)
	}
	if tx.Error != nil {
		return nil
	}
	return &user
}

func GetUserById(id uint, db *gorm.DB) *model.User {
	var user model.User
	tx := db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		slog.Error("[database][GetUserById]user not find :", "id", id, "error", tx.Error.Error())
		return nil
	}
	return &user
}

func RegisterUser(username string, password string, db *gorm.DB) error {
	user := model.User{
		UserName:    username,
		Password:    password,
		DisplayName: "默认名字",
	}

	tx := db.Create(&user)
	err := tx.Error
	if err != nil {
		var mysqlErr *mysql.MySQLError //必须是指针，因为是指针实现了error接口
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 { //违反uniq key
				return logger.Error("[database][RegisterUser] fail , repeat register", "username", username, "error", mysqlErr.Error())
			}
		}
		return logger.Error("[database][RegisterUser]", "username", username, "error", mysqlErr.Error())
	}
	return nil
}

func UpdateUserSingleField(uid uint, field, oldVal, newVal string, db *gorm.DB, needMatchOldValToDB bool) error {

	updater := NewUserUpdater(uid, db, field, oldVal, newVal)

	if err := updater.Prepare(); err != nil {
		return err
	}

	if err := updater.CheckEqual(); err != nil {
		return err
	}

	if needMatchOldValToDB {
		if err := updater.CheckIsEqualToDatabase(); err != nil {
			return err
		}
	}

	if err := updater.Update(); err != nil {
		return err
	}

	slog.Info("[database][UpdatePassword] success", "id", uid, "field", field, "oldVal", oldVal, "newVal", newVal)
	return nil
}
