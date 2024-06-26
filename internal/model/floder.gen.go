// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFloder = "floder"

// Floder mapped from table <floder>
type Floder struct {
	ID        int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"-"`
	Name      string         `gorm:"column:name;type:longtext;comment:文件夹的名字叫啥" json:"name"`                 // 文件夹的名字叫啥
	UserID    string         `gorm:"column:user_id;type:varchar(191);not null;comment:属于哪个用户" json:"userId"` // 属于哪个用户
	Decknum   int64          `gorm:"column:decknum;type:bigint;not null;comment:有几个牌组" json:"decknum"`       // 有几个牌组
}

// TableName Floder's table name
func (*Floder) TableName() string {
	return TableNameFloder
}
