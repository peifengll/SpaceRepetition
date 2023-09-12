package model

import (
	"gorm.io/gorm"
)

// 外边调用方法用的
var DeckNsp Deck

type Deck struct {
	gorm.Model          // 这里边有个更新时间
	Name         string `gorm:"column:name"`         //名字叫什么
	Cardnum      int    `gorm:"column:cardnum"`      //总共有多少
	LearnNumber  int    `gorm:"column:learnnumber"`  //每天的学习量
	Introduction string `gorm:"column:introduction"` //这东西的简介
	FloderId     int    `gorm:"column:floderid"`     //属于哪一个文件夹
	//Version    uint   `gorm:"column:version"`
	//UserId     uint   `gorm:"column:userid"` //

}

func (d *Deck) TableName() string {
	return "decks"
}

// PersonTableName 私人的表
func (d *Deck) PersonTableName(userid string) string {
	return userid + "_" + d.TableName()
}

// AddDeck 增加一个牌组
func (d *Deck) AddDeck(db *gorm.DB, tableName string, deck *Deck) error {
	err := db.Table(tableName).Create(deck).Error
	return err
}

// DeleteDeckByID 根据id删除一条数据
func (d *Deck) DeleteDeckByID(db *gorm.DB, tableName string, id uint) error {
	err := db.Table(tableName).Where("ID = ?", id).Delete(&Deck{}).Error
	return err
}

// GetAllDecks 查 寻所有
func (d *Deck) GetAllDecks(db *gorm.DB, tableName string) ([]*Deck, error) {
	var deckList = make([]*Deck, 0)
	err := db.Table(tableName).Find(&deckList).Error
	if err != nil {
		return nil, err
	}
	return deckList, nil
}

// 按照floderid 查decks
func (d *Deck) GetDecksByFloder(db *gorm.DB, tableName string, floderid int) ([]*Deck, error) {
	var deckList = make([]*Deck, 0)
	err := db.Table(tableName).
		Where("floderid = ?", floderid).
		Find(&deckList).Error

	if err != nil {
		return nil, err
	}
	return deckList, nil
}

// UpdateDeckNameByID 可以改个名字，
func (d *Deck) UpdateDeckNameByID(db *gorm.DB, deckName, tableName string, id uint) error {
	var dic = map[string]interface{}{
		"introduction": deckName,
	}
	err := db.Table(tableName).Where("ID = ?", id).Updates(dic).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateDeckIntrByID 可以改个简介，
func (d *Deck) UpdateDeckIntrByID(db *gorm.DB, intro, tableName string, id uint) error {
	var dic = map[string]interface{}{
		"name": intro,
	}
	err := db.Table(tableName).Where("ID = ?", id).Updates(dic).Error
	if err != nil {
		return err
	}
	return nil
}
