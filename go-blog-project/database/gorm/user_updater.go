// user.go
package database

import (
	"errors"
	"reflect"
	"sdbh/database/model"

	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

type UserUpdater struct {
	uid    uint
	user   *model.User
	field  string
	oldVal interface{}
	newVal interface{}
	db     *gorm.DB
}

// NewUserUpdater 创建更新器
func NewUserUpdater(uid uint, db *gorm.DB, field string, oldVal, newVal interface{}) *UserUpdater {
	return &UserUpdater{
		uid:    uid,
		db:     db,
		field:  field,
		oldVal: oldVal,
		newVal: newVal,
		user:   nil,
	}
}

// Prepare 单条SQL完成：查询字段值 + 判断用户是否存在
func (u *UserUpdater) Prepare() error {
	// 初始化一个 User 结构体实例
	user := model.User{}

	// 执行单条SQL：查询指定用户的指定字段
	// SQL等价于：SELECT `field` FROM `users` WHERE id = ? LIMIT 1
	result := u.db.Model(&model.User{}).
		Select(u.field).
		Where("id = ?", u.uid).
		First(&user)

	err := result.Error
	// 处理错误：用户不存在或查询失败
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Info("[UserUpdater] user not found", "uid", u.uid, "field", u.field)
			return errors.New("用户不存在")
		}
		slog.Error("[UserUpdater] query failed", "uid", u.uid, "field", u.field, "error", err)
		return errors.New("查询字段失败")
	}
	u.user = &user
	return nil
}

// CheckEqual 检查新旧值是否相同
func (u *UserUpdater) CheckEqual() error {
	if reflect.DeepEqual(u.oldVal, u.newVal) {
		slog.Info("[UserUpdater] new value equals old value", "uid", u.uid, "field", u.field)
		return errors.New("新值与旧值相同")
	}
	return nil
}

// Update 执行数据库更新
func (u *UserUpdater) Update() error {
	tx := u.db.Debug().Model(u.user).Where("id = ?", u.uid).Update(u.field, u.newVal)
	if tx.Error != nil {
		slog.Error("[UserUpdater] update failed", "uid", u.uid, "field", u.field, "error", tx.Error)
		return errors.New("更新失败")
	}
	slog.Info("[UserUpdater] update success", "uid", u.uid, "field", u.field)
	return nil
}

func (u *UserUpdater) CheckIsEqualToDatabase() error {
	var err error
	switch u.field {
	case "name":
		err = CheckStringValueIsEqual(u.user.UserName, u.oldVal.(string))
	case "display_name":
		err = CheckStringValueIsEqual(u.user.DisplayName, u.oldVal.(string))
	case "password":
		err = CheckStringValueIsEqual(u.user.Password, u.oldVal.(string))
	}

	if err != nil {
		slog.Info("[UserUpdater] new value not equal to database value", "uid", u.uid, "field", u.field)
		return err
	}
	return nil
}

func CheckStringValueIsEqual(oldVal, newVal string) error {
	if oldVal != newVal {
		return errors.New("旧值与数据库值不匹配")
	}
	return nil
}
