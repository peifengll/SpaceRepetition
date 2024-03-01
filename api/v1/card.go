package v1

import "github.com/peifengll/SpaceRepetition/pkg/helper/fsrs"

//import "github.com/open-spaced-repetition/go-fsrs"

// Knowledge mapped from table <knowledge>
type CardResp struct {
	ID         int64   `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	Font       string  `gorm:"column:font;type:longtext" json:"font"`
	Originfont string  `gorm:"column:originfont;type:longtext" json:"originfont"`
	Back       string  `gorm:"column:back;type:longtext" json:"back"`
	Onlearning int64   `gorm:"column:onlearning;type:tinyint(1)" json:"onlearning"`
	Typeof     int64   `gorm:"column:typeof;type:bigint" json:"typeof"`
	Deckid     int64   `gorm:"column:deckid;type:bigint unsigned" json:"deckid"`
	Skilled    float64 `gorm:"column:skilled;type:double" json:"skilled"`
	Sectionid  int64   `gorm:"column:sectionid;type:bigint" json:"sectionid"`
}
type CardRequest struct {
	ID         int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Font       string `gorm:"column:font;type:longtext" json:"font"`
	Originfont string `gorm:"column:originfont;type:longtext" json:"originfont"`
	Back       string `gorm:"column:back;type:longtext" json:"back"`
	Typeof     int64  `gorm:"column:typeof;type:bigint" json:"typeof"`
	Deckid     int64  `gorm:"column:deckid;type:bigint unsigned" json:"deckid"`
	Sectionid  int64  `gorm:"column:sectionid;type:bigint" json:"sectionid"`
}

type CardIdReq struct {
	Ids []int64 `json:"ids"`
}

type CardReviewResp struct {
	ID         int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	RecordID   int64  `json:"record_id,string"`
	Font       string `gorm:"column:font;type:longtext" json:"font"`
	Originfont string `gorm:"column:originfont;type:longtext" json:"originfont"`
	Back       string `gorm:"column:back;type:longtext" json:"back"`
}

type CardReviewOptReq struct {
	RecordID int64 `json:"record_id,string"`
	ID       int64 `json:"id,string"`
	Opt      fsrs.Rating
}
