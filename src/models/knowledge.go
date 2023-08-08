package models

import (
	"gorm.io/gorm"
)

type Knowledge struct {
	gorm.Model
	Font    string
	Back    string
	NoteId  uint
	Version int //同步版本号
	DeckId  uint
	UserId  uint
}

func (k *Knowledge) TableName() string {
	return "Knowledge"
}
