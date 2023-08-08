package models

import "gorm.io/gorm"

type Note struct {
	Id     uint
	Text   string
	UserId uint
	gorm.DeletedAt
}

func (n *Note) TableName() string {
	return "note"
}
