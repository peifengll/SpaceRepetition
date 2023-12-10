package repository

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
)

type FloderRepository interface {
	FirstById(id int64) (*model.Floder, error)
}

func NewFloderRepository(repository *Repository) FloderRepository {
	return &floderRepository{
		Repository: repository,
	}
}

type floderRepository struct {
	*Repository
}

func (r *floderRepository) FirstById(id int64) (*model.Floder, error) {
	var floder model.Floder
	// TODO: query db
	return &floder, nil
}
