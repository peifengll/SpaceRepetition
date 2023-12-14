package service

import (
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/peifengll/SpaceRepetition/internal/repository"
)

type SectionService interface {
	GetSection(id int64) (*model.Section, error)
	AddSection(v *v1.SectionReq, id string) error
	DeleteSectionById(id int64) error
	UpdateSectionName(id int64, name string) error
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

func (s *sectionService) AddSection(v *v1.SectionReq, id string) error {
	ss := &model.Section{
		Deckid: v.Deckid,
		Name:   v.Name,
		UserID: id,
	}
	return s.query.Section.Create(ss)
}

func (s *sectionService) DeleteSectionById(id int64) error {
	_, err := s.query.Section.Where(s.query.Section.ID.Eq(id)).Delete()
	return err
}

func (s *sectionService) UpdateSectionName(id int64, name string) error {
	u := s.query.Section
	_, err := u.Where(u.ID.Eq(id)).Update(u.Name, name)
	return err
}
