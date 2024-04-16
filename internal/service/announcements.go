package service

import (
	"context"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/repository"
)

type AnnouncementsService interface {
	GetAnnouncements(ctx context.Context, id int64) (*model.Announcement, error)
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

func (s *announcementsService) GetAnnouncements(ctx context.Context, id int64) (*model.Announcement, error) {
	return s.announcementsRepository.GetAnnouncements(ctx, id)
}
