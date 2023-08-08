package models

import "gorm.io/gorm"

type Deck struct {
	gorm.Model
	Name    string
	UserId  uint
	Version uint
}

func (d *Deck) TableName() string {
	return "deck"
}
