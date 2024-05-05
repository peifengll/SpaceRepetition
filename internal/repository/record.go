package repository

import (
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"time"
)

type RecordRepository interface {
	FirstById(id int64) (*model.Record, error)
	QueryDataByTimeRange(userid string, start, end time.Time) ([]*v1.RevInfo, error)
}

func NewRecordRepository(repository *Repository) RecordRepository {
	return &recordRepository{
		Repository: repository,
	}
}

type recordRepository struct {
	*Repository
}

func (r *recordRepository) FirstById(id int64) (*model.Record, error) {
	var record model.Record
	// TODO: query db
	return &record, nil
}

func (r *recordRepository) QueryDataByTimeRange(userid string, start, end time.Time) ([]*v1.RevInfo, error) {
	var records []*v1.RevInfo
	sql := `
	select knowledge_id as card_id, FLOOR(UNIX_TIMESTAMP(LastReview)*1000) as review_time, rate as review_rating,State as review_state
	from record
	where user_id = ?
	  and LastReview between ? and ?
	`
	err := r.db.Raw(sql, userid, start, end).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

/**
select knowledge_id as card_id, FLOOR(UNIX_TIMESTAMP(LastReview)*1000) as review_time, lastop as review_rating
from record
where user_id = 'BLNOXd36vf'
  and LastReview between '2022-05-12' and '2024-05-06'
*/
