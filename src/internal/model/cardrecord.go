package model

import (
	"SpaceRepetition/src/utils/go-fsrs-main"
	"gorm.io/gorm"
	"time"
)

var RecordNsp Record

// 也就是所谓的调度信息
type Record struct {
	gorm.Model
	KnowledgeId uint `json:"knowledgeid"` // 绑定一个knowledge

	Due           time.Time  `gorm:"column:Due"`
	Stability     float64    `gorm:"column:Stability"`
	Difficulty    float64    `gorm:"column:Difficulty"`
	ElapsedDays   uint64     `gorm:"column:ElapsedDays"`
	ScheduledDays uint64     `gorm:"column:ScheduledDays"`
	Reps          uint64     `gorm:"column:Reps"`
	Lapses        uint64     `gorm:"column:Lapses"`
	State         fsrs.State `gorm:"column:State"`
	On            bool       `gorm:"column:On"`     //是否被使用
	LastOp        int        `gorm:"column:lastop"` // 上一次的选择
	LastReview    time.Time  `gorm:"column:LastReview"`
}

func (r *Record) TableName() string {
	return "record"
}
func (r *Record) PersonTableName(userid string) string {
	return userid + "_record"
}

// 增 加一条记录
func (r *Record) AddNewRecord(db *gorm.DB, tableName string, re *Record) error {
	err := db.Table(tableName).Create(&re).Error
	if err != nil {
		return err
	}
	return nil
}

// 按照id撤销最新的一条记录
func (r *Record) RepealNewRecordById(db *gorm.DB, tableName string, id int) error {
	err := db.Table(tableName).
		Where("id = ?", id).
		Delete(&Record{}).
		Error
	if err != nil {
		return err
	}
	return nil
}

// 查出最新的一条数据，
func (r *Record) RepealNewRecord(db *gorm.DB, tableName string, id int) (*Record, error) {
	temp := &Record{}
	err := db.Table(tableName).
		Order("UpdatedAt desc").
		Limit(1).
		First(&temp).
		Error
	if err != nil {
		return nil, err
	}
	return temp, nil
}

// RepealOughtReview 查出所有 在今天之前到期的的  要复习的record 肯定是按照日期查询，
func (r *Record) RepealOughtReview(db *gorm.DB, tableName string) ([]*Record, error) {
	reLi := make([]*Record, 0)
	err := db.Table(tableName).Where("Due <= now").Find(&reLi).Error
	if err != nil {
		return nil, err
	}
	return reLi, nil
}

// RepealRecordByKid 根据kid查询出最新的那一条， on=false  && kid =id
func (r *Record) RepealRecordByKid(db *gorm.DB, tableName string, kid int) (*Record, error) {
	re := &Record{}
	err := db.Table(tableName).Where("Knowledge_Id = ? and `on` = 0 ", kid).First(&re).Error
	if err != nil {
		return nil, err
	}
	return re, nil
}

// 改

// UpdateOnById 根据id查询出出来，on 改为true, 代表已经使用了
func (r *Record) UpdateOnById(db *gorm.DB, tableName string, id int) error {
	err := db.Table(tableName).
		Where("id = ?", id).
		Update("on", 1).
		Error
	if err != nil {
		return err
	}
	return nil
}
