package repository

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
)

type SectionRepository interface {
	FirstById(id int64) (*model.Section, error)
}

func NewSectionRepository(repository *Repository) SectionRepository {
	return &sectionRepository{
		Repository: repository,
	}
}

type sectionRepository struct {
	*Repository
}

func (r *sectionRepository) FirstById(id int64) (*model.Section, error) {
	var section model.Section
	// TODO: query db
	return &section, nil
}
