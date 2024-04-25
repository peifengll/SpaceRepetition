package service

import (
	"errors"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/peifengll/SpaceRepetition/internal/repository"
)

type FloderService interface {
	GetFloder(id int64) (*model.Floder, error)
	FindByUserId(UserID string) ([]*model.Floder, error)
	AddFloder(userid string, name string) error
	DeleteFloder(id int64) error
	CheckAccess(id int64, userid string) (bool, error)
	GetFloderById(iid int64) ([]v1.FloderDeckResp, error)
	UpdateFloder(floder *model.Floder) error
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
func (s *floderService) FindByUserId(UserID string) ([]*model.Floder, error) {
	return s.query.Floder.Where(s.query.Floder.UserID.Eq(UserID)).Find()
}
func (s *floderService) AddFloder(userid string, name string) error {
	f := model.Floder{UserID: userid, Name: name}
	return s.query.Floder.Create(&f)
}

func (s *floderService) DeleteFloder(id int64) error {
	tx := s.query.Begin()
	decks, err := tx.Deck.Where(tx.Deck.Floderid.Eq(id)).Find()
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, v := range decks {
		err = deleteDeckAndCard(v.ID, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	_, err = s.query.Floder.Where(tx.Floder.ID.Eq(id)).Delete()
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (s *floderService) CheckAccess(id int64, userid string) (bool, error) {
	first, err := s.query.Floder.Where(s.query.Floder.ID.Eq(id)).First()
	if err != nil {
		return false, err
	}

	if first.UserID != userid {
		return false, nil
	}
	return true, nil
}

func (s *floderService) GetFloderById(iid int64) ([]v1.FloderDeckResp, error) {
	return nil, errors.New("not implemented")
}

func (s *floderService) UpdateFloder(floder *model.Floder) error {
	_, err := s.query.Floder.Where(s.query.Floder.ID.Eq(floder.ID)).Updates(floder)
	if err != nil {
		return err
	}
	return nil
}
