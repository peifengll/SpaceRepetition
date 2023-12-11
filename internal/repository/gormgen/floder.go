package gormgen

import (
	"gorm.io/gen"
)

type FloderRepositoryG interface {
	// FIndByUserId 查出属于某一个用户的id

	// sql(select * from @@table where user_id = @UserID)
	FindByUserId(UserID string) ([]gen.T, error)
	//sql(delete from @@table where id = @id)
	DelOneByID(id int64) error
}
