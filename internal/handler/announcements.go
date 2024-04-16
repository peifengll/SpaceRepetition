package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/peifengll/SpaceRepetition/internal/service"
)

type AnnouncementsHandler struct {
	*Handler
	announcementsService service.AnnouncementsService
}

func NewAnnouncementsHandler(
    handler *Handler,
    announcementsService service.AnnouncementsService,
) *AnnouncementsHandler {
	return &AnnouncementsHandler{
		Handler:      handler,
		announcementsService: announcementsService,
	}
}

func (h *AnnouncementsHandler) GetAnnouncements(ctx *gin.Context) {

}
