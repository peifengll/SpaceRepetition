package repository

import (
	"context"
	"github.com/peifengll/SpaceRepetition/internal/global_"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"strconv"
)

type AnnouncementsRepository interface {
	GetAnnouncements(ctx context.Context) ([]model.Announcement, error)
	Create(ctx context.Context, ann *model.Announcement) error
	Del(ctx context.Context, id int64) error
	Update(ctx context.Context, info *model.Announcement) error
	GetAnnouncementByIds(ctx context.Context, aids []int64) ([]model.Announcement, error)
	GetAnnouncementIds(ctx context.Context) ([]int64, error)
}

func NewAnnouncementsRepository(
	repository *Repository,
) AnnouncementsRepository {
	return &announcementsRepository{
		Repository: repository,
	}
}

type announcementsRepository struct {
	*Repository
}

func (r *announcementsRepository) GetAnnouncements(ctx context.Context) ([]model.Announcement, error) {
	var announcements []model.Announcement
	err := r.DB(ctx).Raw("select * from announcements").Find(&announcements).Error
	return announcements, err
}

// 有了新的公告， 向redis里插入一个新的Id  然后维持大小为10
func (r *announcementsRepository) Create(ctx context.Context, ann *model.Announcement) error {
	err := r.DB(ctx).Create(ann).Error
	if err != nil {
		return err
	}
	err = r.rdb.RPush(ctx, global_.GetAnnouncementToReadKey(), ann.ID).Err()
	if err != nil {
		return err
	}
	num, err := r.rdb.LLen(ctx, global_.GetAnnouncementToReadKey()).Result()
	if err != nil {
		return err
	}
	err = r.rdb.LPopCount(ctx, global_.GetAnnouncementToReadKey(), max(int(num)-10, 0)).Err()
	return err
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (r *announcementsRepository) Del(ctx context.Context, id int64) error {
	err := r.DB(ctx).Delete(&model.Announcement{}, id).Error
	return err
}

func (r *announcementsRepository) Update(ctx context.Context, info *model.Announcement) error {
	err := r.DB(ctx).Updates(info).Error
	return err
}

func (r *announcementsRepository) GetAnnouncementByIds(ctx context.Context, aids []int64) ([]model.Announcement, error) {
	var infos []model.Announcement
	// 查询要哪些个 的公告
	err := r.DB(ctx).Model(&model.Announcement{}).Where("id in ?", aids).Find(&infos).Error
	return infos, err
}

// 获取这几天还要查看的公告的id
func (r *announcementsRepository) GetAnnouncementIds(ctx context.Context) ([]int64, error) {
	result, err := r.rdb.LRange(ctx, global_.GetAnnouncementToReadKey(), 0, -1).Result()
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	var ids []int64
	for _, v := range result {
		parseInt, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, parseInt)
	}
	return ids, nil
}
