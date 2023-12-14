package v1

type DeckResp struct {
	ID           int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	Name         string `gorm:"column:name;type:longtext" json:"name"`
	Cardnum      int64  `gorm:"column:cardnum;type:bigint" json:"cardnum"`
	Learnnumber  int64  `gorm:"column:learnnumber;type:bigint" json:"learnnumber"`
	Introduction string `gorm:"column:introduction;type:longtext" json:"introduction"`
	Floderid     int64  `gorm:"column:floderid;type:bigint" json:"floderid"`
}

type DeckRequest struct {
	ID           int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name         string `gorm:"column:name;type:longtext" json:"name"`
	Introduction string `gorm:"column:introduction;type:longtext" json:"introduction"`
	Floderid     int64  `gorm:"column:floderid;type:bigint" json:"floderid"`
}

type DeckCardResp struct {
	ID           int64              `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name         string             `gorm:"column:name;type:longtext" json:"name"`
	Introduction string             `gorm:"column:introduction;type:longtext" json:"introduction"`
	Sections     []*SectionCardResp `json:"sections"`
}

type DeckCardReviewResp struct {
	ID    int64             `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name  string            `gorm:"column:name;type:longtext" json:"name"`
	Cards []*CardReviewResp `json:"cards"`
}
