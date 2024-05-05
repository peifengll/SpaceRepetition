package v1

import "time"

type ExportRevlogCsvReq struct {
	Start time.Time
	End   time.Time
	Date1 int64 `json:"date1"`
	Date2 int64 `json:"date2"`
}

// 导出位revlog 的时候用
type RevInfo struct {
	CardId         int   `json:"card_id"  gorm:"column:card_id" csv:"card_id"`
	ReviewTime     int64 `json:"review_time"  gorm:"column:review_time" csv:"review_time"`
	ReviewRating   int   `json:"review_rating"  gom:"column:review_rating" csv:"review_rating"`
	ReviewState    int   `json:"review_state"  gorm:"column:review_state" csv:"review_state"`
	ReviewDuration int   `json:"review_duration"  gorm:"column:review_duration" csv:"review_duration"`
}

// 日期是一个 用户id也是一个
type ExportTask struct {
	Start    time.Time
	End      time.Time
	UserId   string
	ExportId int64 // 如果创建成功的话
}
