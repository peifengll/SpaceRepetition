package repository

import (
	"context"
	"github.com/peifengll/SpaceRepetition/internal/model"
)

type AnnouncementReadRecordRepository interface {
	GetAnnouncementReadRecord(ctx context.Context, id int64) (*model.AnnouncementReadRecord, error)
	AddAnnouncementReadRecord(ctx context.Context, info *model.AnnouncementReadRecord) error
	GetAnnouncementReadRecordIDByIds(ctx context.Context, userid string, ids []int64) ([]int64, error)
}

func NewAnnouncementReadRecordRepository(
	repository *Repository,
) AnnouncementReadRecordRepository {
	return &announcementReadRecordRepository{
		Repository: repository,
	}
}

type announcementReadRecordRepository struct {
	*Repository
}

func (r *announcementReadRecordRepository) GetAnnouncementReadRecord(ctx context.Context, id int64) (*model.AnnouncementReadRecord, error) {
	var announcementReadRecord model.AnnouncementReadRecord

	return &announcementReadRecord, nil
}
func (r *announcementReadRecordRepository) AddAnnouncementReadRecord(ctx context.Context, info *model.AnnouncementReadRecord) error {
	return r.DB(ctx).Create(info).Error
}

// 获取没读的公告的id
func (r *announcementReadRecordRepository) GetAnnouncementReadRecordIDByIds(ctx context.Context, userid string, ids []int64) ([]int64, error) {
	var infos []model.AnnouncementReadRecord
	err := r.DB(ctx).Model(&model.AnnouncementReadRecord{}).Where("user_id= ? and announcement_id in ?", userid, ids).Select("announcement_id").Find(&infos).Error
	if err != nil {
		return nil, err
	}
	ma := make(map[int64]bool)
	for i := 0; i < len(infos); i++ {
		ma[infos[i].AnnouncementID] = true
	}
	var data []int64
	for _, v := range ids {
		if !ma[v] {
			data = append(data, v)
		}
	}
	return data, err
}
