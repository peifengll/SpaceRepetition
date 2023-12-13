package service

import (
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/peifengll/SpaceRepetition/internal/repository"
)

type KnowledgeService interface {
	GetKnowledge(id int64) (*model.Knowledge, error)
	AddCard(v *v1.CardRequest, userid string) error
	GetKnowledgeById(id int64) (*v1.CardResp, error)
	UpdateCard(v *v1.CardRequest) error
	DeleteCard(id int64) error
	SearchCards(content string) ([]*v1.CardResp, error)
	ChooseToReview(ids []int) error
}

func NewKnowledgeService(que *query.Query, service *Service, knowledgeRepository repository.KnowledgeRepository) KnowledgeService {
	return &knowledgeService{
		Service: service,
		query:   que,

		knowledgeRepository: knowledgeRepository,
	}
}

type knowledgeService struct {
	*Service
	query               *query.Query
	knowledgeRepository repository.KnowledgeRepository
}

func (s *knowledgeService) GetKnowledge(id int64) (*model.Knowledge, error) {
	return s.knowledgeRepository.FirstById(id)
}

func (s *knowledgeService) AddCard(v *v1.CardRequest, userid string) error {
	cc := &model.Knowledge{
		Font:       v.Font,
		Originfont: v.Originfont,
		Back:       v.Back,
		Typeof:     v.Typeof,
		Deckid:     v.Deckid,
		Sectionid:  v.Sectionid,
		UserID:     userid,
	}
	return s.query.Knowledge.Create(cc)
}

func (s *knowledgeService) GetKnowledgeById(id int64) (*v1.CardResp, error) {
	first, err := s.query.Knowledge.Where(s.query.Knowledge.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	res := &v1.CardResp{
		ID:         first.ID,
		Font:       first.Back,
		Originfont: first.Originfont,
		Back:       first.Back,
		Onlearning: first.Onlearning,
		Typeof:     first.Typeof,
		Deckid:     first.Deckid,
		Skilled:    first.Skilled,
		Sectionid:  first.Sectionid,
	}
	return res, nil
}

func (s *knowledgeService) UpdateCard(v *v1.CardRequest) error {
	k := &model.Knowledge{
		Font:       v.Font,
		Originfont: v.Originfont,
		Back:       v.Back,
		Typeof:     v.Typeof,
		Deckid:     v.Deckid,
		Sectionid:  v.Sectionid,
	}
	_, err := s.query.Knowledge.Where(s.query.Knowledge.ID.Eq(v.ID)).Updates(k)
	if err != nil {
		return err
	}
	return nil
}

func (s *knowledgeService) DeleteCard(id int64) error {
	_, err := s.query.Knowledge.Where(s.query.Knowledge.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *knowledgeService) SearchCards(content string) ([]*v1.CardResp, error) {
	// todo 详细的查询规则
	knows, err := s.query.Knowledge.Where().Find()
	if err != nil {
		return nil, err
	}
	cards := make([]*v1.CardResp, len(knows))
	for i := 0; i < len(cards); i++ {
		cards[i] = &v1.CardResp{
			ID:         knows[i].ID,
			Font:       knows[i].Font,
			Originfont: knows[i].Originfont,
			Back:       knows[i].Back,
			Onlearning: knows[i].Onlearning,
			Typeof:     knows[i].Typeof,
			Deckid:     knows[i].Deckid,
			Skilled:    knows[i].Skilled,
			Sectionid:  knows[i].Sectionid,
		}
	}
	return cards, nil
}

func (s *knowledgeService) ChooseToReview(ids []int) error {
	//	 todo  这里要用事务来做，全部都开始复习
	panic("not implement")
}
