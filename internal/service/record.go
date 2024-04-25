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

// 根据card删除 记录
// todo  这中要排除报错为 没有那个数据的情况
func delRecordTx(kid int64, tx *query.QueryTx) error {
	_, err := tx.Record.Where(tx.Record.KnowledgeID.Eq(kid)).Delete()
	return err
}
