package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"net/http"
	"strconv"
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
		Handler:              handler,
		announcementsService: announcementsService,
	}
}

func (h *AnnouncementsHandler) GetAnnouncements(ctx *gin.Context) {

}

func (h *AnnouncementsHandler) AddAnnouncement(ctx *gin.Context) {
	var req v1.AnnouncementAddReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	id, err := strconv.Atoi(GetUserIdFromCtx(ctx))
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	req.AdminID = int64(id)
	err = h.announcementsService.AddAnnouncement(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *AnnouncementsHandler) ShowAnnouncement(ctx *gin.Context) {
	res, err := h.announcementsService.GetAnnouncements(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, res)
}

func (h *AnnouncementsHandler) DelAnnouncement(ctx *gin.Context) {

	var req v1.AnnouncementDelReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	//id, err := strconv.Atoi(GetUserIdFromCtx(ctx))
	//if err != nil {
	//	v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
	//	return
	//}

	err := h.announcementsService.DelAnnouncement(ctx, req.ID)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *AnnouncementsHandler) UpdateAnnouncement(ctx *gin.Context) {
	var req v1.AnnouncementUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	err := h.announcementsService.UpdateAnnouncement(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}
