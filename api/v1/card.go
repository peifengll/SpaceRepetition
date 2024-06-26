package v1

import (
	"github.com/peifengll/SpaceRepetition/pkg/helper/fsrs"
	"time"
)

//import "github.com/open-spaced-repetition/go-fsrs"

// Knowledge mapped from table <knowledge>
type CardResp struct {
	ID         int64     `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	Font       string    `gorm:"column:font;type:longtext" json:"font"`
	Originfont string    `gorm:"column:originfont;type:longtext" json:"originfont"`
	Back       string    `gorm:"column:back;type:longtext" json:"back"`
	Onlearning int64     `gorm:"column:onlearning;type:tinyint(1)" json:"onlearning"`
	Typeof     int64     `gorm:"column:typeof;type:bigint" json:"typeof"`
	Deckid     int64     `gorm:"column:deckid;type:bigint unsigned" json:"deckid"`
	Skilled    float64   `gorm:"column:skilled;type:double" json:"skilled"`
	Sectionid  int64     `gorm:"column:sectionid;type:bigint" json:"sectionid"`
	Due        time.Time `gorm:"column:Due;type:datetime(3);not null;comment:到期时间，也就是该复习的日子" json:"due"` // 到期时间，也就是该复习的日子
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
	Id int64 `json:"id,int"`
}

type CardReviewResp struct {
	ID         int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	RecordID   int64  `gorm:"column:record_id" json:"record_id"`
	Font       string `gorm:"column:font;type:longtext" json:"font"`
	Originfont string `gorm:"column:originfont;type:longtext" json:"originfont"`
	Back       string `gorm:"column:back;type:longtext" json:"back"`
	DeckID     int64  `gorm:"column:deckid" json:"deckid"`
	Typeof     int64  `gorm:"column:typeof;type:bigint" json:"typeof"`
}

type CardReviewOptReq struct {
	RecordID int64       `json:"record_id"`
	ID       int64       `json:"id"`
	Opt      fsrs.Rating `json:"opt"`
}
