package v1

type FloderDeckResp struct {
	ID      int64      `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	Name    string     `gorm:"column:name;type:longtext;comment:文件夹的名字叫啥" json:"name"`  // 文件夹的名字叫啥
	Decknum int64      `gorm:"column:decknum;type:bigint;comment:有几个牌组" json:"decknum"` // 有几个牌组
	Decks   []DeckResp `json:"decks"`
}

type GetAllFloderAndDecksResponse struct {
}
