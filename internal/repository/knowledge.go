package repository

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
)

type KnowledgeRepository interface {
	FirstById(id int64) (*model.Knowledge, error)
}

func NewKnowledgeRepository(repository *Repository) KnowledgeRepository {
	return &knowledgeRepository{
		Repository: repository,
	}
}

type knowledgeRepository struct {
	*Repository
}

func (r *knowledgeRepository) FirstById(id int64) (*model.Knowledge, error) {
	var knowledge model.Knowledge
	// TODO: query db
	return &knowledge, nil
}
