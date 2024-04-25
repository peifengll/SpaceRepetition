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
	tx := s.query.Begin()
	cards, err := tx.Knowledge.Where(tx.Knowledge.Sectionid.Eq(id)).Find()
	if err != nil {
		tx.Rollback()
		return err
	}
	section, err := tx.Section.Where(tx.Section.ID.Eq(id)).First()
	if err != nil {
		tx.Rollback()
		return err
	}
	deckid := section.Deckid
	// 删除卡片 ，更改数量
	num := len(cards)
	// 更改数量
	_, err = tx.Deck.Where(tx.Deck.ID.Eq(deckid)).UpdateSimple(tx.Deck.Cardnum.Sub(int64(num)))
	if err != nil {
		tx.Rollback()
		return err
	}
	// 删除复习记录
	for _, v := range cards {
		err = delRecordTx(v.ID, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	_, err = tx.Knowledge.Where(tx.Knowledge.Sectionid.Eq(id)).Delete()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Section.Where(s.query.Section.ID.Eq(id)).Delete()
	if err != nil {
		tx.Rollback()
		return err
	}
	//1.  删除章节下的卡片
	//2.  更改deck的数量
	//3.  删除卡片的记录
	return tx.Commit()
}

func (s *sectionService) UpdateSectionName(id int64, name string) error {
	u := s.query.Section
	_, err := u.Where(u.ID.Eq(id)).Update(u.Name, name)
	return err
}
