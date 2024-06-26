// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameExportInfo = "export_info"

// ExportInfo 用户导出复习记录的情况（复习记录用于fsrs weight优化）
type ExportInfo struct {
	ID            int64     `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	UserID        string    `gorm:"column:user_id;type:varchar(50);not null" json:"userId"`
	RevlogPath    string    `gorm:"column:revlog_path;type:varchar(255);not null;comment:导出的revlog的路径" json:"revlogPath"` // 导出的revlog的路径
	TrainDataPath string    `gorm:"column:train_data_path;type:varchar(255);comment:训练结果的路径" json:"trainDataPath"`        // 训练结果的路径
	DataStartTime time.Time `gorm:"column:data_start_time;type:date;not null;comment:复习数据的开始时间" json:"dataStartTime"`     // 复习数据的开始时间
	DataEndTime   time.Time `gorm:"column:data_end_time;type:date;not null;comment:复习数据的结束时间" json:"dataEndTime"`         // 复习数据的结束时间
	ExportDate    time.Time `gorm:"column:export_date;type:date;not null;comment:导出revlog的日期" json:"exportDate"`          // 导出revlog的日期
	Columns_      int64     `gorm:"column:columns;type:int;comment:导出的revlog有多少列" json:"columns"`                         // 导出的revlog有多少列
	State         int64     `gorm:"column:state;type:int;comment:0已完成，1等待中（不少等待就是完成就好）" json:"state"`                     // 0已完成，1等待中（不少等待就是完成就好）
}

// TableName ExportInfo's table name
func (*ExportInfo) TableName() string {
	return TableNameExportInfo
}
