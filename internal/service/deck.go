package service

import (
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/peifengll/SpaceRepetition/internal/repository"
)

type DeckService interface {
	GetDeckAll(id int64) (*v1.DeckCardResp, error)
	GetDecksByFloderId(id int64) ([]*model.Deck, error)
	CheckAddDeckAccess(floderid int64, id string) (bool, error)
	AddDeck(de *model.Deck) error
	CheckDeckAccess(iid int64, id string) (bool, error)
	UpdateDeck(deck *model.Deck) error
	DeleteDeck(id int64) error
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

func (s *deckService) GetDeckAll(id int64) (*v1.DeckCardResp, error) {
	first, err := s.query.Deck.Where(s.query.Deck.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	res := &v1.DeckCardResp{
		ID:           id,
		Name:         first.Introduction,
		Introduction: first.Introduction,
	}
	// 查询所有属于这个deck的section
	secs, err := s.query.Section.Where(s.query.Section.Deckid.Eq(id)).Find()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(secs); i++ {
		//	查询 所有 属于这个section的card
		cardmodels, err := s.query.Knowledge.Where(s.query.Knowledge.Sectionid.Eq(secs[i].ID)).Find()
		if err != nil {
			return nil, err
		}
		cards := make([]*v1.CardResp, len(cardmodels))
		for i := 0; i < len(cardmodels); i++ {
			cards[i] = &v1.CardResp{
				ID:         cardmodels[i].ID,
				Font:       cardmodels[i].Font,
				Originfont: cardmodels[i].Originfont,
				Back:       cardmodels[i].Back,
				Onlearning: cardmodels[i].Onlearning,
				Typeof:     cardmodels[i].Typeof,
				Deckid:     cardmodels[i].Deckid,
				Skilled:    cardmodels[i].Skilled,
				Sectionid:  cardmodels[i].Sectionid,
			}
		}
		res.Sections = append(res.Sections, &v1.SectionCardResp{
			ID:     secs[i].ID,
			Deckid: id,
			Name:   secs[i].Name,
			Cards:  cards,
		})
	}

	return res, nil
}

func (s *deckService) GetDecksByFloderId(id int64) ([]*model.Deck, error) {
	first, err := s.query.Deck.Where(s.query.Deck.Floderid.Eq(id)).Find()
	if err != nil {
		return nil, err
	}
	return first, nil
}

func (s *deckService) CheckAddDeckAccess(floderid int64, id string) (bool, error) {
	first, err := s.query.Floder.Where(s.query.Floder.ID.Eq(floderid)).First()
	if err != nil {
		return false, err
	}

	if first.UserID != id {
		return false, nil
	}
	return true, nil
}

func (s *deckService) AddDeck(de *model.Deck) error {
	tx := s.query.Begin()
	err := tx.Deck.Create(de)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Floder.Where(tx.Floder.ID.Eq(de.Floderid)).UpdateSimple(tx.Floder.Decknum.Add(1))
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *deckService) CheckDeckAccess(iid int64, id string) (bool, error) {
	first, err := s.query.Deck.Where(s.query.Deck.ID.Eq(iid)).First()
	if err != nil {
		return false, err
	}

	if first.UserID != id {
		return false, nil
	}
	return true, nil
}

func (s *deckService) UpdateDeck(deck *model.Deck) error {
	_, err := s.query.Deck.Where(s.query.Deck.ID.Eq(deck.ID)).Updates(deck)
	if err != nil {
		return err
	}
	return nil
}

func (s *deckService) DeleteDeck(id int64) error {
	tx := s.query.Begin()
	err := deleteDeckAndCard(id, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// 删除下边所有卡片,然后删除牌组
func deleteDeckAndCard(id int64, tx *query.QueryTx) error {
	first, err := tx.Deck.Where(tx.Deck.ID.Eq(id)).First()
	if err != nil {
		return err
	}
	flowerId := first.Floderid
	_, err = tx.Knowledge.Where(tx.Knowledge.Deckid.Eq(id)).Delete()
	if err != nil {
		return err
	}
	_, err = tx.Section.Where(tx.Section.Deckid.Eq(id)).Delete()
	if err != nil {
		return err
	}
	_, err = tx.Deck.Where(tx.Deck.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	// 牌组删掉了，数量减一个
	_, err = tx.Floder.Where(tx.Floder.ID.Eq(flowerId)).UpdateSimple(tx.Floder.Decknum.Sub(1))
	return err
}
