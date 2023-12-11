package service

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/peifengll/SpaceRepetition/internal/repository"
)

type DeckService interface {
	GetDeck(id int64) (*model.Deck, error)
	GetDecksByFloderId(id int64) ([]*model.Deck, error)
}

func NewDeckService(que *query.Query, service *Service, deckRepository repository.DeckRepository) DeckService {
	return &deckService{
		Service: service,
		query:   que,

		deckRepository: deckRepository,
	}
}

type deckService struct {
	*Service
	query          *query.Query
	deckRepository repository.DeckRepository
}

func (s *deckService) GetDeck(id int64) (*model.Deck, error) {
	return s.deckRepository.FirstById(id)
}

func (s *deckService) GetDecksByFloderId(id int64) ([]*model.Deck, error) {
	first, err := s.query.Deck.Where(s.query.Deck.Floderid.Eq(id)).Find()
	if err != nil {
		return nil, err
	}
	return first, nil
}
