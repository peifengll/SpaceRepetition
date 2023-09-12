package model

import "gorm.io/gorm"

var KnowledgeNsp Knowledge

type Knowledge struct {
	gorm.Model
	Font       string  `gorm:"column:font"`
	OriginFont string  `gorm:"column:originfont"` // 原本，备份用的
	Back       string  `gorm:"column:back"`
	OnLearning bool    `gorm:"column:onlearning"` //是否在学
	Typeof     int     `gorm:"column:typeof"`     // 是问答型，还是挖空型
	DeckId     uint    `gorm:"column:deckid"`     // 属于哪个牌组,自然不用考虑属于哪个文件夹
	Skilled    float64 `gorm:"column:skilled"`    // 熟练度
	SectionId  int     `gorm:"column:sectionid"`
	//NoteId  uint 这个不要了
	//Version int //同步版本号
	//UserId uint
}

func (k *Knowledge) TableName() string {
	return "Knowledge"
}

// PersonTableName 私人的表
func (k *Knowledge) PersonTableName(userid string) string {
	return userid + "_" + k.TableName()
}

// 增
func (k *Knowledge) AddKnowledge(db *gorm.DB, tableName string, kn *Knowledge) error {
	return db.Table(tableName).Create(kn).Error
}

// 删
func (k *Knowledge) DeleteKnowledgeById(db *gorm.DB, tableName string, id uint) error {
	return db.Table(tableName).Where("ID = ?", id).Delete(&Knowledge{}).Error
}

// 查 一条
func (k *Knowledge) FindKnowledgeById(db *gorm.DB, tableName string, kid int) (*Knowledge, error) {
	var res *Knowledge
	err := db.Table(tableName).Where("id = ?", kid).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 查 一条
func (k *Knowledge) FindKnowledgeIfOnLearningById(db *gorm.DB, tableName string, kid int) (bool, error) {
	var res *Knowledge
	err := db.Table(tableName).Where("id = ?", kid).First(&res).Error
	if err != nil {
		return false, err
	}
	return res.OnLearning, nil
}

// 查所有
func (k *Knowledge) FindAllKnowledge(db *gorm.DB, tableName string) ([]*Knowledge, error) {
	res := make([]*Knowledge, 0)
	err := db.Table(tableName).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FindAllKnowledgeOngoing true是在学的，false是没有在学的 按照ongoing查
func (k *Knowledge) FindAllKnowledgeOngoing(db *gorm.DB, tableName string, on bool) ([]*Knowledge, error) {
	res := make([]*Knowledge, 0)
	err := db.Table(tableName).Where("ongoing = ?", on).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 按照sectionId 查
func (k *Knowledge) FindKnowLedgeListBySectionId(db *gorm.DB, tableName string, sectionid int) ([]*Knowledge, error) {
	kl := make([]*Knowledge, 0)
	err := db.Table(tableName).
		Where("sectionid = ?", sectionid).
		Find(&kl).
		Error
	if err != nil {
		return nil, err
	}
	return kl, nil
}

// FindKnowledgeOrderBySkilled 在学的，按照熟练度排序
func (k *Knowledge) FindKnowledgeOrderBySkilled(db *gorm.DB, tableName string) ([]*Knowledge, error) {
	res := make([]*Knowledge, 0)
	err := db.Table(tableName).
		Where("ongoing = ?", true).
		Order("skilled").
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// todo 修改更新操作
func (k *Knowledge) UpdateOnLearning(db *gorm.DB, tableName string, kid int) error {
	//先看是不是已经在了，
	is, err := k.FindKnowledgeIfOnLearningById(db, tableName, kid)
	if err != nil {
		return err
	}
	if is == true {
		//不用调度
		return nil
	}
	//	 加入调度
	err = db.Table(tableName).
		Where("id = ?", kid).
		Update("onlearning", 1).
		Error
	if err != nil {
		return err
	}
	return nil
}

// 复原
