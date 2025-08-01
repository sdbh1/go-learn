package model

type User struct {
	Id          int    `gorm:"primaryKey" `
	UserName    string `gorm:"column:name" `
	DisplayName string `gorm:"column:display_name"`
	Password    string `gorm:"column:password"`
	PostNum     int    `gorm:"column:post_num" default:"0"`
}
