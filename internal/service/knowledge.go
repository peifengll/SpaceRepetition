package service

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/peifengll/SpaceRepetition/internal/repository"
)

type KnowledgeService interface {
	GetKnowledge(id int64) (*model.Knowledge, error)
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
