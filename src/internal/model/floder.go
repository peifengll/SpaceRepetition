package model

import (
	"gorm.io/gorm"
)

var FloderNsp Floder

type Floder struct {
	gorm.Model
	Name    string `gorm:"column:name"`
	DeckNum int    `gorm:"column:decknum"`
}

func (f *Floder) TableName() string {
	return "floder"
}

func (f *Floder) PersonTableName(userid string) string {
	return userid + "_floder"
}

// 增
func (f *Floder) CreateFloder(db *gorm.DB, tableName, fname string) error {
	var c Floder
	c.Name = fname
	c.DeckNum = 0
	err := db.Table(tableName).Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

// 删
// 查 返回所有
func (f *Floder) GetAllFloder(db *gorm.DB, tableName string) ([]*Floder, error) {
	fl := make([]*Floder, 0)
	err := db.Table(tableName).Find(&fl).Error
	if err != nil {
		return nil, err
	}
	return fl, nil
}

//改

func (f *Floder) UpdateFloderNameById(db *gorm.DB, tableName, fname string, id int) error {
	dic := map[string]any{
		"name": fname,
	}
	err := db.Table(tableName).
		Where("id = ?", id).
		Updates(dic).
		Error
	if err != nil {
		return err
	}
	return nil
}
