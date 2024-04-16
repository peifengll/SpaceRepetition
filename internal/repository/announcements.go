package repository

import (
	"context"
	"github.com/peifengll/SpaceRepetition/internal/model"
)

type AnnouncementsRepository interface {
	GetAnnouncements(ctx context.Context) ([]model.Announcement, error)
	Create(ctx context.Context, ann *model.Announcement) error
	Del(ctx context.Context, id int64) error
	Update(ctx context.Context, info *model.Announcement) error
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

func (r *announcementsRepository) Create(ctx context.Context, ann *model.Announcement) error {
	err := r.DB(ctx).Create(ann).Error
	return err
}

func (r *announcementsRepository) Del(ctx context.Context, id int64) error {
	err := r.DB(ctx).Delete(&model.Announcement{}, id).Error
	return err
}

func (r *announcementsRepository) Update(ctx context.Context, info *model.Announcement) error {
	err := r.DB(ctx).Updates(info).Error
	return err
}
