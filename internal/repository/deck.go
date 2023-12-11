package repository

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
)

type DeckRepository interface {
	FirstById(id int64) (*model.Deck, error)
}

func NewDeckRepository(repository *Repository) DeckRepository {
	return &deckRepository{
		Repository: repository,
	}
}

type deckRepository struct {
	*Repository
}

func (r *deckRepository) FirstById(id int64) (*model.Deck, error) {
	var deck model.Deck
	// TODO: query db
	return &deck, nil
}
