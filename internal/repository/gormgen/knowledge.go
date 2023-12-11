package gormgen

import "gorm.io/gen"

// 这个其实就是对应一张卡片的实体

type KnowledgeRepositoryG interface {
	//	按照deckid 查询出这个deckid下的所有卡片

	//  sql(select * from @@table where deckid = @deckId)
	FindByDeckId(deckId int64) ([]gen.T, error)

	//	 更新原本的内容,font and back  但是这个应该属于基础实现...?

	// 将选中的ids进行复习

	//sql(update @@table set onlearning = 1 where id in (@ids))
	SetOnLearning(ids []int) error
}
