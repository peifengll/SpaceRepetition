package service

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/peifengll/SpaceRepetition/internal/repository"
)

type SectionService interface {
	GetSection(id int64) (*model.Section, error)
}

func NewSectionService(que *query.Query, service *Service, sectionRepository repository.SectionRepository) SectionService {
	return &sectionService{
		Service: service,
		query:   que,

		sectionRepository: sectionRepository,
	}
}

type sectionService struct {
	*Service
	query             *query.Query
	sectionRepository repository.SectionRepository
}

func (s *sectionService) GetSection(id int64) (*model.Section, error) {
	return s.sectionRepository.FirstById(id)
}
