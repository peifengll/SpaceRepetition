// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID               int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	CreatedAt        time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"` // 创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"` // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"-"`
	Gender           int64          `gorm:"column:gender;type:bigint;comment:性别" json:"gender"` // 性别
	Age              int64          `gorm:"column:age;type:bigint" json:"age"`
	Password         string         `gorm:"column:password;type:longtext;not null" json:"password"`
	HeadURL          string         `gorm:"column:head_url;type:longtext;comment:头像" json:"headUrl"` // 头像
	Email            string         `gorm:"column:email;type:longtext;not null" json:"email"`
	Username         string         `gorm:"column:username;type:longtext;comment:登录的用户名" json:"username"` // 登录的用户名
	Introduction     string         `gorm:"column:introduction;type:longtext" json:"introduction"`
	Learnnumoneday   int64          `gorm:"column:learnnumoneday;type:bigint;comment:一天学多少" json:"learnnumoneday"` // 一天学多少
	UserID           string         `gorm:"column:user_id;type:varchar(191);not null" json:"userId"`
	Nickname         string         `gorm:"column:nickname;type:varchar(20)" json:"nickname"`
	MaxInterval      int64          `gorm:"column:max_interval;type:int;comment:最大复习间隔，超过这个就算复习完成（推荐365） 默认365" json:"maxInterval"`                  // 最大复习间隔，超过这个就算复习完成（推荐365） 默认365
	Weights          string         `gorm:"column:weights;type:varchar(80);comment:复习计算需要的参数" json:"weights"`                                        // 复习计算需要的参数
	RequestRetention float32        `gorm:"column:request_retention;type:float;comment:想要下次见到一张卡片时，回忆起的概率（推荐70%~90%,默认90%）" json:"requestRetention"` // 想要下次见到一张卡片时，回忆起的概率（推荐70%~90%,默认90%）
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
