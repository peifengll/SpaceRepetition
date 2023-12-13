package v1

type SectionResp struct {
	ID     int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	Deckid int64  `gorm:"column:deckid;type:bigint" json:"deckid"`
	Name   string `gorm:"column:name;type:longtext" json:"name"`
}
type SectionCardResp struct {
	ID     int64       `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	Deckid int64       `gorm:"column:deckid;type:bigint" json:"deckid"`
	Name   string      `gorm:"column:name;type:longtext" json:"name"`
	Cards  []*CardResp `json:"cards"`
}
