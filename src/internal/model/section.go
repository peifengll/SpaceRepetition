package model

import "gorm.io/gorm"

var SectionNsp Section

type Section struct {
	gorm.Model
	// id为0就是默认章节
	DeckId int    `gorm:"column:deckid"`
	Name   string `gorm:"column:name"`
}

func (s *Section) TableName() string {
	return "section"
}

func (s *Section) PersonTableName(userid string) string {
	return userid + "_" + s.TableName()
}

func (s *Section) AddSection(db *gorm.DB, tableName, sname string, deckid int) error {
	se := &Section{
		DeckId: deckid,
		Name:   sname,
	}
	err := db.Table(tableName).Create(&se).Error
	if err != nil {
		return err
	}
	return nil
}

// 按照deckid 查询sectionlist
func (s *Section) FindSectionListByDeckId(db *gorm.DB, tableName string, deckid int) ([]*Section, error) {
	slist := make([]*Section, 0)
	err := db.Table(tableName).Where("deckid = ?", deckid).Find(&slist).Error
	if err != nil {
		return nil, err
	}
	return slist, nil

}
