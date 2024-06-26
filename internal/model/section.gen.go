// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSection = "section"

// Section mapped from table <section>
type Section struct {
	ID        int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"-"`
	Deckid    int64          `gorm:"column:deckid;type:bigint;not null" json:"deckid"`
	Name      string         `gorm:"column:name;type:longtext" json:"name"`
	UserID    string         `gorm:"column:user_id;type:varchar(191);not null;comment:属于哪一个用户" json:"userId"` // 属于哪一个用户
	CardNum   int64          `gorm:"column:card_num;type:int;not null;comment:章节下卡片的数量" json:"cardNum"`       // 章节下卡片的数量
}

// TableName Section's table name
func (*Section) TableName() string {
	return TableNameSection
}
