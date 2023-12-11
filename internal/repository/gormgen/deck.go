package gormgen

import "gorm.io/gen"

type DeckRepositoryG interface {
	// 按照文件id查询出这个下边有多少个牌组，和牌组的情况
	// sql(select * from @@table where floderid = @FloderID)
	FindByFloderID(FloderID int64) ([]gen.T, error)

	//sql(delete from @@table where id = @id)
	DelOneByID(id int64) error
}
