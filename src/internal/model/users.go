package model

import (
	"gorm.io/gorm"
)

var UserNsp User

type User struct {
	gorm.Model
	Gender         int    `gorm:"column:gender"`   //性别
	Age            int    `gorm:"column:age"`      //年龄
	PassWord       string `gorm:"column:password"` //密码
	HeadUrl        string `gorm:"column:head_url"`
	Email          string `gorm:"column:email"`
	UserName       string `gorm:"column:username"`
	Introduction   string `gorm:"column:introduction"`
	LearnNumOneDay int    `gorm:"column:learnnumoneday"`
	//NickName       string `gorm:"column:nickname"` //昵称

}

func (u *User) TableName() string {
	return "user"
}

// 增
func (u *User) AddUser(db *gorm.DB, user *User) error {
	err := db.Model(&User{}).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// 删
// 查 验证账户密码
func (u *User) FindByUsernameAndPassword(db *gorm.DB, username, password string) (*User, error) {
	k := &User{}
	err := db.Model(&User{}).Where("username = ? and password = ?", username, password).First(&k).Error
	//fmt.Println("err  ....  ?", k, err)
	if err != nil {
		return nil, err
	}
	k.PassWord = "********"
	return k, nil
}

// 查 名字重不,true就是重复
func (u *User) FindByUsername(db *gorm.DB, username string) (bool, error) {
	var c int64
	err := db.Model(&User{}).Where("username = ? ", username).Count(&c).Error
	if err != nil {
		return false, err
	}
	if c > 0 {
		return true, nil
	}
	return false, nil
}

// 查 id
func (u *User) FindIdByUsername(db *gorm.DB, username string) (int, error) {
	c := &User{}
	err := db.Model(&User{}).Where("username = ? ", username).First(c).Error
	if err != nil {
		return 0, err
	}
	return int(c.ID), nil
}

// 改
func (u *User) UpdateByUsername(db *gorm.DB, info *User) error {
	err := db.Model(&User{}).Updates(&info).Error
	if err != nil {
		return err
	}
	return nil
}
