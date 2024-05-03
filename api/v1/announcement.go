package v1

import "time"

type AnnouncementAddReq struct {
	AdminID int64  `gorm:"column:admin_id;type:int;not null" json:"adminId"`
	Title   string `gorm:"column:title;type:varchar(100);not null" json:"title"`
	Content string `gorm:"column:content;type:text;not null" json:"content"`
}

type AnnouncementUpdateReq struct {
	ID      int64  `json:"id"`
	Title   string `gorm:"column:title;type:varchar(100);not null" json:"title"`
	Content string `gorm:"column:content;type:text;not null" json:"content"`
}

type AnnouncementDelReq struct {
	ID int64 `json:"id"`
}

type AnnouncementReadReq struct {
	ID int64 `json:"id"`
}

type AnnouncementsUserResp struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishTime time.Time ` json:"publishTime"`
}
