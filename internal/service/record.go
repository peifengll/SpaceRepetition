package service

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/peifengll/SpaceRepetition/internal/repository"
)

type RecordService interface {
	GetRecord(id int64) (*model.Record, error)
}

func NewRecordService(que *query.Query, service *Service, recordRepository repository.RecordRepository) RecordService {
	return &recordService{
		Service: service,
		query:   que,

		recordRepository: recordRepository,
	}
}

type recordService struct {
	*Service
	query            *query.Query
	recordRepository repository.RecordRepository
}

func (s *recordService) GetRecord(id int64) (*model.Record, error) {
	return s.recordRepository.FirstById(id)
}
