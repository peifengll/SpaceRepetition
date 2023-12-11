package repository

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
)

type RecordRepository interface {
	FirstById(id int64) (*model.Record, error)
}

func NewRecordRepository(repository *Repository) RecordRepository {
	return &recordRepository{
		Repository: repository,
	}
}

type recordRepository struct {
	*Repository
}

func (r *recordRepository) FirstById(id int64) (*model.Record, error) {
	var record model.Record
	// TODO: query db
	return &record, nil
}
