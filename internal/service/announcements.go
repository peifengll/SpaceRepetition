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
	SetAnnouncementHaveRead(ctx context.Context, id int64, userid string) error
	GetAnnouncementNoRead(ctx context.Context, userid string) ([]v1.AnnouncementsUserResp, error)
}

func NewAnnouncementsService(service *Service, announcementsRepository repository.AnnouncementsRepository, arecord repository.AnnouncementReadRecordRepository) AnnouncementsService {
	return &announcementsService{
		Service:                 service,
		announcementsRepository: announcementsRepository,
		arecord:                 arecord,
	}
}

type announcementsService struct {
	*Service
	announcementsRepository repository.AnnouncementsRepository
	arecord                 repository.AnnouncementReadRecordRepository
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

func (s *announcementsService) SetAnnouncementHaveRead(ctx context.Context, id int64, userid string) error {
	info := &model.AnnouncementReadRecord{
		UserID:         userid,
		AnnouncementID: id,
	}
	return s.arecord.AddAnnouncementReadRecord(ctx, info)
}

func (s *announcementsService) GetAnnouncementNoRead(ctx context.Context, userid string) ([]v1.AnnouncementsUserResp, error) {
	// 获取最新的几个公告
	ids, err := s.announcementsRepository.GetAnnouncementIds(ctx)
	if err != nil {
		return nil, err
	}
	// 看要通知的是那些
	aids, err := s.arecord.GetAnnouncementReadRecordIDByIds(ctx, userid, ids)
	if err != nil {
		return nil, err
	}
	infos, err := s.announcementsRepository.GetAnnouncementByIds(ctx, aids)
	if err != nil {
		return nil, err
	}
	var data []v1.AnnouncementsUserResp
	for _, v := range infos {
		data = append(data, v1.AnnouncementsUserResp{
			ID:          v.ID,
			Title:       v.Title,
			Content:     v.Content,
			PublishTime: v.PublishTime,
		})
	}
	return data, nil
}
