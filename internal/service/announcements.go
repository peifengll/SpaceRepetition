package service

import (
	"context"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/repository"
	"time"
)

type AnnouncementsService interface {
	GetAnnouncements(ctx context.Context) ([]model.Announcement, error)
	AddAnnouncement(ctx context.Context, rq *v1.AnnouncementAddReq) error
	DelAnnouncement(ctx context.Context, id int64) error
	UpdateAnnouncement(ctx context.Context, rq *v1.AnnouncementUpdateReq) error
}

func NewAnnouncementsService(service *Service, announcementsRepository repository.AnnouncementsRepository) AnnouncementsService {
	return &announcementsService{
		Service:                 service,
		announcementsRepository: announcementsRepository,
	}
}

type announcementsService struct {
	*Service
	announcementsRepository repository.AnnouncementsRepository
}

func (s *announcementsService) GetAnnouncements(ctx context.Context) ([]model.Announcement, error) {
	return s.announcementsRepository.GetAnnouncements(ctx)
}

func (s *announcementsService) AddAnnouncement(ctx context.Context, rq *v1.AnnouncementAddReq) error {
	info := &model.Announcement{
		AdminID:     rq.AdminID,
		Title:       rq.Title,
		Content:     rq.Content,
		PublishTime: time.Now(),
	}
	return s.announcementsRepository.Create(ctx, info)
}

func (s *announcementsService) DelAnnouncement(ctx context.Context, id int64) error {
	return s.announcementsRepository.Del(ctx, id)
}
func (s *announcementsService) UpdateAnnouncement(ctx context.Context, rq *v1.AnnouncementUpdateReq) error {
	info := &model.Announcement{
		ID:      rq.ID,
		Title:   rq.Title,
		Content: rq.Content,
	}
	return s.announcementsRepository.Update(ctx, info)
}
