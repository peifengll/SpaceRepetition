package query

import "gorm.io/gorm"

func NewQuery(db *gorm.DB) *Query {
	return Use(db)
}
