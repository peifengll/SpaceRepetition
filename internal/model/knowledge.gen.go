// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameKnowledge = "knowledge"

// Knowledge mapped from table <knowledge>
type Knowledge struct {
	ID         int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	CreatedAt  time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"-"`
	Font       string         `gorm:"column:font;type:longtext" json:"font"`
	Originfont string         `gorm:"column:originfont;type:longtext" json:"originfont"`
	Back       string         `gorm:"column:back;type:longtext" json:"back"`
	Onlearning int64          `gorm:"column:onlearning;type:tinyint(1)" json:"onlearning"`
	Typeof     int64          `gorm:"column:typeof;type:bigint" json:"typeof"`
	Deckid     int64          `gorm:"column:deckid;type:bigint unsigned" json:"deckid"`
	Skilled    float64        `gorm:"column:skilled;type:double" json:"skilled"`
	Sectionid  int64          `gorm:"column:sectionid;type:bigint" json:"sectionid"`
	UserID     string         `gorm:"column:user_id;type:varchar(191)" json:"userId"`
}

// TableName Knowledge's table name
func (*Knowledge) TableName() string {
	return TableNameKnowledge
}
