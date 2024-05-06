package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"net/http"
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
	//id, err := strconv.Atoi(GetUserIdFromCtx(ctx))
	//if err != nil {
	//	v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
	//	return
	//}
	//req.AdminID = int64(id)
	if req.AdminID == 0 {
		req.AdminID = 2
	}
	err := h.announcementsService.AddAnnouncement(ctx, &req)
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

// 设置该公告为已读，不再推送
func (h *AnnouncementsHandler) SetAnnouncementHaveRead(ctx *gin.Context) {
	var req v1.AnnouncementReadReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	err := h.announcementsService.SetAnnouncementHaveRead(ctx, req.ID, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

// 该用户还没查看的公告
func (h *AnnouncementsHandler) GetAnnouncementNoRead(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	anns, err := h.announcementsService.GetAnnouncementNoRead(ctx, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, anns)
}
