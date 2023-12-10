package service

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/peifengll/SpaceRepetition/internal/repository"
)

type FloderService interface {
	GetFloder(id int64) (*model.Floder, error)
	FindByUserId(UserID int64) ([]model.Floder, error)
}

func NewFloderService(que *query.Query, service *Service, floderRepository repository.FloderRepository) FloderService {
	return &floderService{
		query:            que,
		Service:          service,
		floderRepository: floderRepository,
	}
}

type floderService struct {
	*Service
	query            *query.Query
	floderRepository repository.FloderRepository
}

func (s *floderService) GetFloder(id int64) (*model.Floder, error) {
	return s.floderRepository.FirstById(id)
}
func (s *floderService) FindByUserId(UserID int64) ([]model.Floder, error) {
	return s.query.Floder.FindByUserId(UserID)
}
