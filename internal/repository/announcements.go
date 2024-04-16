package repository

import (
	"context"
	"github.com/peifengll/SpaceRepetition/internal/model"
)

type AnnouncementsRepository interface {
	GetAnnouncements(ctx context.Context, id int64) (*model.Announcement, error)
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

func (r *announcementsRepository) GetAnnouncements(ctx context.Context, id int64) (*model.Announcement, error) {
	var announcements model.Announcement

	return &announcements, nil
}
