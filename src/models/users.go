package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Gender   string `gorm:"column:gender"`   //性别
	Age      int    `gorm:"column:age"`      //年龄
	PassWord string `gorm:"column:password"` //密码
	NickName string `gorm:"column:nickname"` //昵称
	HeadUrl  string `gorm:"column:head_url"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:pwd"`
	UserName string `gorm:"column:username"`
}

func (u *User) TableName() string {
	return "user"
}
