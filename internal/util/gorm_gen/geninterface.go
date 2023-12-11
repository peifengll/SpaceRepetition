package main

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/repository/gormgen"
	"gorm.io/gen"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../query",
		Mode:    gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
	})

	g.ApplyBasic(model.Floder{}, model.Deck{}, model.Knowledge{})

	g.ApplyInterface(func(gormgen.FloderRepositoryG) {}, model.Floder{})

	g.Execute()
}

//
//type FloderRepositoryG interface {
//	// FIndByUserId 查出属于某一个用户的id
//
//	// sql(select * from @@table where user_id = @UserID)
//	FindByUserId(UserID int64) ([]gen.T, error)
//	//sql(delete from @@table where id = @id)
//	DelOneByID(id int64) error
//}
